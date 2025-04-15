"""
Copyright © 2025 LeavesWebber

SPDX-License-Identifier: MPL-2.0

Feel free to contact LeavesWebber@outlook.com

该程序生成随机不重复的 PDF 文件，文件名是随机常用单词，文件内容是 base58 随机 hash 值，文件用于测试上传论文
"""

# 设置控制台输出编码
import sys
import io
import os
import random
import string
import hashlib
import subprocess
import importlib
sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')

# 常用英文单词列表（长度在3-12个字母之间）
COMMON_WORDS = [
    'Apple', 'Book', 'Cloud', 'Data', 'Earth', 'Flower', 'Garden', 'House',
    'Image', 'Jewel', 'Light', 'Music', 'Night', 'Ocean', 'Paper', 'Queen',
    'River', 'Sun', 'Tree', 'Water', 'World', 'Youth', 'Zebra', 'Brain',
    'Chair', 'Dance', 'Eagle', 'Forest', 'Globe', 'Heart', 'Island', 'Jungle',
    'Lemon', 'Mountain', 'Nature', 'Orange', 'Peace', 'Rainbow', 'Star',
    'Universe', 'Valley', 'Window', 'Xylophone', 'Yellow', 'Zenith'
]

def check_and_install_dependencies():
    """检查并安装必要的依赖"""
    dependencies = ['base58', 'reportlab']
    for package in dependencies:
        try:
            importlib.import_module(package)
            print(f"{package} 已安装")
        except ImportError:
            print(f"正在安装 {package}...")
            subprocess.check_call([sys.executable, "-m", "pip", "install", package])
            print(f"{package} 安装完成")

def generate_random_string(length=4):
    """生成指定长度的随机字符串"""
    return ''.join(random.choices(string.ascii_uppercase + string.digits, k=length))

def generate_random_word():
    """生成一个随机的英文单词"""
    return random.choice(COMMON_WORDS)

def generate_random_hash_base58():
    """生成随机的 hash 并转换为 base58 格式"""
    # 导入 base58 模块（确保已安装）
    import base58
    
    random_data = str(random.getrandbits(256)).encode()
    hash_obj = hashlib.sha256(random_data)
    hash_bytes = hash_obj.digest()
    return base58.b58encode(hash_bytes).decode('utf-8')

def create_pdf_with_content(filename, content):
    """创建包含指定内容的 PDF 文件"""
    # 导入 reportlab 模块（确保已安装）
    from reportlab.pdfgen import canvas
    
    c = canvas.Canvas(filename)
    c.setFont("Helvetica", 12)
    
    # 将内容分成多行，避免一行太长
    lines = [content[i:i+80] for i in range(0, len(content), 80)]
    y_position = 800
    for line in lines:
        c.drawString(50, y_position, line)
        y_position -= 15
    
    c.save()

def main():
    print("正在检查并安装必要的依赖...")
    check_and_install_dependencies()
    
    # 获取程序所在目录
    script_dir = os.path.dirname(os.path.abspath(__file__))
    output_dir = os.path.join(script_dir, "生成的随机论文")
    os.makedirs(output_dir, exist_ok=True)
    print(f"\n程序所在目录: {script_dir}")
    print(f"PDF文件将保存在: {output_dir}")
    
    print("\n开始生成PDF文件...")
    # 创建7个PDF文件
    for i in range(1, 8):
        # 生成文件名
        file_num = f"{i:02d}"
        random_word = generate_random_word()
        random_suffix = generate_random_string(4)
        filename = f"{file_num}-Test-Paper-{random_word}-{random_suffix}.pdf"
        file_path = os.path.join(output_dir, filename)
        
        # 生成随机内容
        content = generate_random_hash_base58()
        
        # 创建PDF文件
        create_pdf_with_content(file_path, content)
        print(f"已创建文件: {file_path}")
    
    print("\n所有PDF文件生成完成！")
    print(f"文件位置: {output_dir}")

if __name__ == "__main__":
    try:
        main()
    except Exception as e:
        print(f"程序运行出错: {e}")
        input("按Enter键退出...")