# MPERproxy 技术说明文档

## 1. 简介

`MPERproxy` 是一个专为 MPER 代币设计的透明代理合约，遵循 ERC1967 标准。它支持通过治理机制进行逻辑合约的升级，同时代理合约本身的地址保持不变。该合约集成了 OpenZeppelin 的 `TimelockController`，以实现分阶段、有时间延迟的治理操作，增强了安全性。

**核心特性:**
- **透明代理 (ERC1967):** 用户与代理合约交互，代理合约将调用委托给当前的逻辑合约。
- **不可升级的代理:** `MPERproxy` 合约本身的代码是不可变的。
- **可升级的逻辑合约:** MPER 代币的业务逻辑所在的合约可以被替换。
- **两步所有权转移 (`Ownable2Step`):** 用于安全地转移合约所有权。
- **时间锁治理 (`TimelockController`):** 敏感操作（如合约升级）需要经过预设的时间延迟才能执行，为社区提供了审查和反应的时间。

## 2. 架构

- **代理合约 (`MPERproxy`):** 用户交互的入口，存储逻辑合约地址，并将调用委托给逻辑合约。其所有权在部署后立即转移给内部创建的 `TimelockController` 实例。
- **逻辑合约 (例如 `MPER.sol`):** 实现 MPER 代币的核心功能。可以被升级替换。
- **时间锁控制器 (`TimelockController`):** 控制 `MPERproxy` 的所有权。所有对 `MPERproxy` 的管理操作（如发起升级、更改升级延迟）都必须通过此时间锁合约进行提案和执行。

## 3. 关键组件与状态变量

### 导入的合约
- `@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol`: ERC1967 代理标准实现。
- `@openzeppelin/contracts/access/Ownable2Step.sol`: 两步所有权转移机制。
- `@openzeppelin/contracts/utils/introspection/ERC165.sol`: ERC165 接口检测。
- `@openzeppelin/contracts/utils/StorageSlot.sol`: 用于操作特定存储槽位。
- `@openzeppelin/contracts/governance/TimelockController.sol`: 时间锁治理合约。

### 状态变量
- `timelock (TimelockController public)`: 关联的 `TimelockController` 合约实例地址。`MPERproxy` 的所有者。
- `upgradeDelay (uint256 public)`: 升级提案从被调度到可执行的最小等待时间（秒）。此延迟由 `MPERproxy` 内部的 `proposeUpgrade` 函数在调度到 `TimelockController` 时使用。
- `isProposalActive (mapping(bytes32 => bool) public)`: 跟踪通过 `MPERproxy.proposeUpgrade` 发起的升级提案是否处于活动状态。

### 常量
- `MIN_UPGRADE_DELAY (uint256 public constant)`: `upgradeDelay` 的最小允许值 (默认为 1 天)。
- `MAX_REASON_LENGTH (uint256 public constant)`: 升级原因字符串的最大长度 (默认为 256 字节)。

### 存储槽位
- `_IMPLEMENTATION_INTERFACE_ID_SLOT`: 用于存储期望的逻辑合约接口ID，用于升级后验证。
- `_TIMELOCK_SLOT`: 用于存储 `TimelockController` 的地址。

## 4. 核心函数

### `constructor(address logic, address admin, bytes memory data, uint256 initialDelay)`
- 初始化代理合约。
- **参数:**
    - `logic`: 初始逻辑合约的地址。
    - `admin`: 初始管理员地址。此地址将成为创建的 `TimelockController` 的管理员和初始提议者。
    - `data`: 调用初始逻辑合约的初始化函数所需的数据 (例如，`initialize(uint256)` 的编码调用)。
    - `initialDelay`: `TimelockController` 的最小延迟时间。这也是 `MPERproxy` 的初始 `upgradeDelay`。
- **操作:**
    - 设置初始逻辑合约和初始化数据。
    - 校验参数（逻辑合约地址、admin 地址、`initialDelay`）。
    - （可选）存储逻辑合约的 ERC165 接口ID。
    - 创建一个新的 `TimelockController` 实例：
        - 提议者 (`proposers`): 初始为 `admin`。
        - 执行者 (`executors`): 初始为 `MPERproxy` 合约自身 (`address(this)`)。
        - 管理员 (`admin`): 初始为 `admin`。
    - 将 `MPERproxy` 的所有权通过 `_transferOwnership` 转移给新创建的 `TimelockController` 地址。
    - 设置 `upgradeDelay = initialDelay`。
    - 触发 `Initialized` 事件。

### `setUpgradeDelay(uint256 newDelay) external onlyOwner`
- 允许 `MPERproxy` 的所有者（即 `TimelockController`）更改 `upgradeDelay`。
- **注意:** 此操作本身也需要通过 `TimelockController` 的提案-执行流程来完成。
- 触发 `UpgradeDelayChanged` 事件。

### `implementation() external view returns (address)`
- 返回当前逻辑合约的地址。

## 5. 合约升级流程

合约升级是一个涉及 `TimelockController` 的多阶段过程，旨在提供透明度和安全性。

**参与者角色 (在 `TimelockController` 中定义):**
- **Proposer:** 有权向 `TimelockController` 提交操作提案。初始时，这是部署 `MPERproxy` 时指定的 `admin`。
- **Executor:** 有权执行 `TimelockController` 中已通过延迟期的提案。初始时，`MPERproxy` 合约自身 (`address(this)`) 是一个执行者，允许它执行由 `proposeUpgrade` 调度的 `upgradeToAndCall`。
- **Admin:** `TimelockController` 的管理员，可以管理角色等。初始时，这是部署 `MPERproxy` 时指定的 `admin`。

**升级流程概览:**

合约升级通常涉及两个主要的时间锁控制阶段：
1.  **第一阶段 (可选，但推荐):** 通过 `TimelockController` 提案并执行对 `MPERproxy.proposeUpgrade()` 的调用。
2.  **第二阶段:** `MPERproxy.proposeUpgrade()` 函数内部再次使用 `TimelockController` 来调度实际的 `MPERproxy.upgradeToAndCall()` 函数，并应用 `MPERproxy.upgradeDelay`。

**详细步骤:**

**步骤 A: 准备升级 (链下)**
1.  **开发和测试新的逻辑合约版本。**
2.  **准备初始化调用数据 (如果需要):** 如果新的逻辑合约版本包含 `initialize` 函数或需要在升级时调用的其他设置函数，请准备其 ABI 编码的调用数据。

**步骤 B: 通过 `TimelockController` 提案调用 `MPERproxy.proposeUpgrade` (第一阶段时间锁)**

这一步由具有 `PROPOSER_ROLE` 的地址（例如初始 `admin`）在 `TimelockController` 合约上操作。

1.  **调用 `TimelockController.schedule()` 函数:**
    *   `target`: `MPERproxy` 合约的地址。
    *   `value`: 0。
    *   `data`: `MPERproxy.proposeUpgrade(address newImplementation, bytes memory data, string memory reason)` 函数的 ABI 编码调用。
        *   `newImplementation`: 新逻辑合约的地址。
        *   `data`: 用于新逻辑合约初始化调用的数据 (例如 `MPER.initialize(initialSupply)` 的编码)。
        *   `reason`: 升级的原因。
    *   `predecessor`: 通常为 `bytes32(0)`。
    *   `salt`: 一个随机的 `bytes32` 值，用于区分提案。
    *   `delay`: 此操作在 `TimelockController` 中的延迟时间（应大于等于 `TimelockController` 的 `minDelay`，即部署时设置的 `initialDelay`）。

2.  **等待 `TimelockController` 的延迟期结束。**

3.  **调用 `TimelockController.execute()` 函数:**
    *   参数与 `schedule` 时相同 (`target`, `value`, `data`, `predecessor`, `salt`)。
    *   此调用会使 `TimelockController` 合约作为 `msg.sender` 来调用 `MPERproxy.proposeUpgrade()`。

**步骤 C: `MPERproxy.proposeUpgrade()` 执行 (第二阶段时间锁的开始)**

当 `TimelockController` 调用 `MPERproxy.proposeUpgrade()` 时：

1.  **`onlyOwner` 检查:** `MPERproxy` 的 `onlyOwner` 修饰符会验证 `msg.sender` 是否为 `address(timelock)`，此检查通过。
2.  **参数验证:**
    *   `newImplementation` 必须是合约地址。
    *   `reason` 长度必须在有效范围内。
3.  **数据验证 (`_validateUpgradeData`):**
    *   如果 `data` 是对 `initialize(uint256)` 的调用，会检查 `initialSupply > 0`。可以扩展此函数以进行更具体的验证。
4.  **计算 ETA (Estimated Time of Arrival):** `eta = block.timestamp + upgradeDelay`。这里的 `upgradeDelay` 是 `MPERproxy` 合约中定义的延迟。
5.  **准备升级调用:** `upgradeCall = abi.encodeWithSignature("upgradeToAndCall(address,bytes,string)", newImplementation, data, reason)`。
6.  **在 `TimelockController` 上调度实际升级:**
    *   调用 `timelock.schedule(address(this), 0, upgradeCall, predecessor, proposalSalt, eta)`。
        *   `target`: `MPERproxy` 自身 (`address(this)`).
        *   `value`: 0.
        *   `data`: 上一步准备的 `upgradeCall`。
        *   `predecessor`: `bytes32(0)`.
        *   `proposalSalt`: 基于 `block.timestamp` 生成的盐。
        *   `delay` (隐式): `eta - block.timestamp`，即 `MPERproxy` 的 `upgradeDelay`。
7.  **记录提案:** `isProposalActive[proposalId] = true`。`proposalId` 是根据上述参数计算的哈希。
8.  **返回 `proposalId`。**

**步骤 D: 等待 `MPERproxy.upgradeDelay` 结束**

这是第二阶段的时间锁。

**步骤 E: 执行实际升级 (`MPERproxy.upgradeToAndCall`)**

在 `eta` 时间到达后，任何人都可以调用 `TimelockController.execute()` 来执行在步骤 C.6 中调度的操作。

1.  **调用 `TimelockController.execute()`:**
    *   `target`: `MPERproxy` 地址 (`address(this)`).
    *   `value`: 0.
    *   `data`: `upgradeCall` (即 `abi.encodeWithSignature("upgradeToAndCall(address,bytes,string)", newImplementation, data, reason)`).
    *   `predecessor`: `bytes32(0)`.
    *   `salt`: 步骤 C.6 中使用的 `proposalSalt`.
2.  `TimelockController` 会调用 `MPERproxy.upgradeToAndCall(newImplementation, data, reason)`。
3.  **`MPERproxy.upgradeToAndCall()` 内部:**
    *   **权限检查:** `require(msg.sender == address(timelock), "Unauthorized")`。此检查通过。
    *   **执行升级:** `ERC1967Utils.upgradeToAndCall(newImplementation, data)`。这会：
        *   将代理的实现槽更新为 `newImplementation`。
        *   如果 `data` 非空，则使用 `delegatecall` 在新实现合约的上下文中执行 `data`。
    *   **验证新实现 (`_verifyImplementation`):**
        *   检查 `newImplementation` 是否是合约。
        *   如果 `_IMPLEMENTATION_INTERFACE_ID_SLOT` 中设置了接口ID，则验证新实现是否支持该接口 (ERC165)。
    *   **触发事件:** `emit Upgraded(newImplementation, block.timestamp, reason)`。

### 取消升级提案 (`cancelUpgrade(bytes32 proposalId) external onlyOwner`)

此函数允许 `MPERproxy` 的所有者（即 `TimelockController`，因此也需要通过 Timelock 提案-执行流程）取消一个已通过 `proposeUpgrade` 提交但尚未执行的升级提案。

1.  **调用 `TimelockController.schedule()` (或直接由 Admin 调用 `TimelockController.cancel` 如果权限允许):**
    *   `target`: `MPERproxy` 地址。
    *   `data`: `MPERproxy.cancelUpgrade(bytes32 proposalId)` 的 ABI 编码调用。
    *   ... 其他 `schedule` 参数 ...
2.  等待 Timelock 延迟后，`TimelockController.execute()` 调用 `MPERproxy.cancelUpgrade()`.
3.  **`MPERproxy.cancelUpgrade()` 内部:**
    *   `onlyOwner` 检查通过。
    *   `require(isProposalActive[proposalId], "Proposal is not active")`。
    *   调用 `timelock.cancel(proposalId)` 来取消在 `TimelockController` 中对应的调度。
    *   `delete isProposalActive[proposalId]`。
    *   触发 `UpgradeCancelled` 事件。

## 6. 注意事项与安全考量

- **双重时间锁:** 理解升级过程涉及两个潜在的时间锁延迟：
    1.  `TimelockController` 自身的 `minDelay` (在部署 `MPERproxy` 时作为 `initialDelay` 设置)，用于执行对 `MPERproxy.proposeUpgrade` 或 `MPERproxy.setUpgradeDelay` 等管理函数的调用。
    2.  `MPERproxy` 内部的 `upgradeDelay`，用于 `proposeUpgrade` 函数调度实际的 `upgradeToAndCall`。
- **`TimelockController` 管理:**
    - `TimelockController` 的 `admin` 角色（初始为部署 `MPERproxy` 时的 `admin`）拥有最高权限，可以管理角色（提议者、执行者）。务必妥善保管此 `admin` 账户的私钥。
    - 提议者 (`PROPOSER_ROLE`) 可以发起提案。
- **初始化数据 (`data`):**
    - 在 `proposeUpgrade` 时提供的 `data` 参数用于在新逻辑合约升级后立即调用其初始化函数 (例如 `initialize(...)`)。
    - 必须确保 `data` 被正确 ABI 编码。错误的 `data` 可能导致初始化失败或合约状态错误。
    - `_validateUpgradeData` 函数提供了对 `initialize(uint256)` 的基本检查，但对于其他初始化函数签名，需要确保数据正确性。
- **接口验证 (`_verifyImplementation`):**
    - 如果在 `MPERproxy` 的 `_IMPLEMENTATION_INTERFACE_ID_SLOT` 中设置了特定的接口ID，升级时会检查新逻辑合约是否实现了该接口。这有助于防止升级到不兼容的合约。
- **升级原因 (`reason`):**
    - `reason` 字符串会被记录在 `Upgraded` 事件中，应清晰说明升级的目的。长度受 `MAX_REASON_LENGTH` 限制。
- **Gas 成本:** 部署新的逻辑合约和执行升级相关的交易（特别是 `schedule` 和 `execute`）可能会消耗大量 gas。
- **提案ID (`proposalId`):**
    - `proposalId` 是根据提案的各项参数（目标、值、调用数据、前置提案、盐、ETA）计算出的哈希值。它用于在 `TimelockController` 中唯一标识一个操作，并用于取消操作。
    - `MPERproxy` 中的 `isProposalActive` 也使用此 ID。
- **测试:**
    - 在主网部署或升级前，务必在测试网络上对新的逻辑合约和整个升级流程进行彻底测试。
    - 测试应包括成功升级、带初始化数据的升级、以及取消升级等场景。
- **`isProposalActive` 清理:** 当前 `MPERproxy` 的 `upgradeToAndCall` 函数在成功升级后不会清除 `isProposalActive` 中的对应提案状态。虽然 `TimelockController` 在执行后会处理其内部状态，但 `MPERproxy` 中的这个映射状态会保留。这通常无害，因为 `TimelockController` 不会允许同一提案ID被重复执行，但可以考虑在 `upgradeToAndCall` 成功后也清理 `isProposalActive` 以保持状态一致性。
- **安全性:** 代理合约和时间锁机制的安全性高度依赖于 `TimelockController` 管理员账户和提议者账户的安全性。

## 7. 事件

- `Initialized(address indexed implementation, address indexed admin)`: 代理合约初始化时触发。
- `Upgraded(address indexed newImplementation, uint256 timestamp, string reason)`: 逻辑合约成功升级后触发。
- `UpgradeCancelled(bytes32 indexed proposalId)`: 升级提案被取消时触发。
- `UpgradeDelayChanged(uint256 oldDelay, uint256 newDelay)`: `upgradeDelay` 被修改时触发。

## 附录 A: TimelockController 详解

`TimelockController` 是 OpenZeppelin 提供的一个标准合约，用于在执行敏感操作前强制执行一个时间延迟。这为治理参与者提供了审查和反应的时间，从而增强了系统的安全性。在 `MPERproxy` 中，`TimelockController` 实例拥有代理合约的所有权，并控制所有管理操作，如合约升级。

### A.1 核心角色

`TimelockController` 定义了几个关键角色，这些角色通过 `bytes32` 类型的哈希值表示：

-   **`TIMELOCK_ADMIN_ROLE`**: 管理员角色。拥有此角色的账户可以授予和撤销其他角色，以及更改时间锁本身的参数（如最小延迟）。在 `MPERproxy` 的构造函数中，此角色初始授予部署时指定的 `admin` 地址。
-   **`PROPOSER_ROLE`**: 提议者角色。拥有此角色的账户可以向时间锁提交新的操作提案（即调用 `schedule` 函数）。在 `MPERproxy` 的构造函数中，此角色初始授予部署时指定的 `admin` 地址。
-   **`EXECUTOR_ROLE`**: 执行者角色。拥有此角色的账户可以执行已经通过时间延迟且准备就绪的操作提案（即调用 `execute` 函数）。在 `MPERproxy` 的构造函数中，`address(this)` (即 `MPERproxy` 合约自身) 和 `address(0)` (表示任何人都可以执行已就绪的提案) 通常被授予此角色。
-   **`CANCELLER_ROLE`**: 取消者角色。拥有此角色的账户可以取消一个已调度但尚未执行的操作提案（即调用 `cancel` 函数）。通常，`PROPOSER_ROLE` 也兼具取消其自己提案的能力，或者 `TIMELOCK_ADMIN_ROLE` 可以取消任何提案。在 `MPERproxy` 的构造函数中，此角色初始授予部署时指定的 `admin` 地址。
### A.2 核心函数

`TimelockController` 提供了几个核心函数，用于管理提案、执行提案、取消提案以及更改时间锁参数。

-   **`schedule(address target, uint256 value, bytes calldata data, bytes32 predecessor, bytes32 salt, uint256 delay)`**: 此函数允许账户（拥有 `PROPOSER_ROLE`）向时间锁提交一个新的操作提案。提案会在 `delay` 秒后执行。
-   **`execute(address target, uint256 value, bytes calldata data, bytes32 predecessor, bytes32 salt, uint256 delay)`**: 此函数允许账户（拥有 `EXECUTOR_ROLE`）执行一个已经通过时间延迟且准备就绪的操作提案。
-   **`cancel(bytes32 proposalId)`**: 此函数允许账户（拥有 `CANCELLER_ROLE`）取消一个已经调度但尚未执行的操作提案。
-   **`setMinDelay(uint256 minDelay)`**: 此函数允许账户（拥有 `TIMELOCK_ADMIN_ROLE`）更改时间锁的最小延迟时间。


