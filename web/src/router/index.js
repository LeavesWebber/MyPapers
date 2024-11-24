import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/home/Home.vue'
import User from '../views/home/User.vue'
import Test from '../views/home/Test.vue'
import Main from '../views/home/Main.vue'
import Register from '../views/home/Register.vue'
import Login from '../views/home/Login.vue'
import Center from '../views/center/Center.vue'
import Information from '../views/center/Information.vue'
import CreateConference from '../views/center/CreateConference.vue'
import DetailCommittee from '../views/center/DetailCommittee.vue'
import DetailConference from '../views/center/DetailConference.vue'
import DetailJournal from '../views/center/DetailJournal.vue'
import UpdateCommittee from '../views/center/UpdateCommittee.vue'
import UpdateInformation from '../views/center/UpdateInformation.vue'
import UpdateJournal from '../views/center/UpdateJournal.vue'
import UpdateConference from '../views/center/UpdateConference.vue'
import DetailPaper from '../views/center/DetailPaper.vue'
import UpdatePaper from '../views/center/UpdatePaper.vue'
import RevisePaper from '../views/center/RevisePaper.vue'
import ReviewPaperDetail from '../views/center/ReviewPaperDetail.vue'
import JournalIssuesCenter from '../views/center/JournalIssues.vue'
import ConferenceIssuesCenter from '../views/center/ConferenceIssues.vue'
import CreateCommittee from '@/views/center/CreateCommittee.vue'
import Committees from '../views/home/Committees.vue'
import HomeDetailCommittee from '../views/home/DetailCommittee.vue'
import Conferences from '../views/home/Conferences.vue'
import HomeDetailConference from '../views/home/DetailConference.vue'
import Journals from '../views/home/Journals.vue'
import Market from '../views/home/Market.vue'
import HomeDetailJournal from '../views/home/DetailJournal.vue'
import ConferenceSubmit from '../views/home/ConferenceSubmit.vue'
import JournalSubmit from '../views/home/JournalSubmit.vue'
import JournalIssues from '../views/home/JournalIssues.vue'
import ConferenceIssues from '../views/home/ConferenceIssues.vue'
import Publications from '../views/home/Publications.vue'
import HomeDetailPaper from '../views/home/DetailPaper.vue'
import MyNFTs from '../views/center/MyNFTs.vue'
import MyCommittee from '../views/center/SelfCommittee.vue'
import MyConference from '../views/center/SelfConference.vue'
import NFTSelling from '../views/center/NFTSelling.vue'
import { component } from 'vue'
import SelfCommittee from '../views/center/SelfCommittee.vue'
import SelfConference from '../views/center/SelfConference.vue'
import InReviewPapers from '@/views/center/InReviewPapers.vue'
import InReview from '@/views/center/InReview.vue'
import Reviewed from '@/views/center/Reviewed.vue'
import ReviewedPapers from '@/views/center/ReviewedPapers.vue'
Vue.use(VueRouter)
const routes = [
    // 主路由
    {
        path: '/',
        redirect: "/home", // 重定向
        component: Main,
        children: [
            // 子路由 
            { path: 'home', name: 'home', component: Home }, // 首页
            { path: 'committees', name: 'committees', component: Committees },
            { path: 'detailCommittee', component: HomeDetailCommittee },
            { path: 'conferences', name: 'conferences', component: Conferences },
            { path: 'detailConference', component: HomeDetailConference },
            { path: 'journals', name: 'journals', component: Journals },
            { path: 'market', name: 'market', component: Market },
            { path: 'detailJournal', component: HomeDetailJournal },
            { path: 'conferenceSubmit', component: ConferenceSubmit },
            { path: 'journalSubmit', component: JournalSubmit },
            { path: 'journalIssues', component: JournalIssues },
            { path: 'conferenceIssues', component: ConferenceIssues },
            { path: 'publications', component: Publications },
            { path: 'detailPaper', component: HomeDetailPaper },
            { path: 'user', component: User }, // 用户管理
            { path: 'test', component: Test }, // 用户管理
            // { path: 'mall', name: 'mall', component: Mall }, // 商品管理
            // { path: 'page1', name: 'page1', component: PageOne }, // 页面1
            // { path: 'Page2', name: 'page2', component: PageTwo } // 页面2
        ]
    },
    {
        path: '/register',
        name: 'register',
        component: Register
    },
    {
        path: '/login',
        name: 'login',
        component: Login,
    },
    {
        path: '/center',
        component: Center,
        name: 'Center',
        redirect: "/center/information", // 重定向
        meta: { requireAuth: true },
        children: [
            // 子路由 
            { path: '/center/detailCommittee', component: DetailCommittee, meta: { requireAuth: true } },
            { path: '/center/detailConference', component: DetailConference, meta: { requireAuth: true } },
            { path: '/center/detailJournal', component: DetailJournal, meta: { requireAuth: true } },
            { path: '/center/updateInformation', component: UpdateInformation, meta: { requireAuth: true } },
            { path: '/center/updateCommittee', component: UpdateCommittee, meta: { requireAuth: true } },
            { path: '/center/updateJournal', component: UpdateJournal, meta: { requireAuth: true } },
            { path: '/center/updateConference', component: UpdateConference, meta: { requireAuth: true } },
            { path: '/center/detailPaper', name: DetailPaper, component: DetailPaper, meta: { requireAuth: true } },
            { path: '/center/updatePaper', name: UpdatePaper, component: UpdatePaper, meta: { requireAuth: true } },
            { path: '/center/revisePaper', name: RevisePaper, component: RevisePaper, meta: { requireAuth: true } },
            { path: '/center/reviewPaperDetail', name: ReviewPaperDetail, component: ReviewPaperDetail, meta: { requireAuth: true } },
            { path: '/center/journalIssues', name: JournalIssuesCenter, component: JournalIssuesCenter, meta: { requireAuth: true } },
            { path: '/center/conferenceIssues', name: ConferenceIssuesCenter, component: ConferenceIssuesCenter, meta: { requireAuth: true } },
            { path: '/center/myNFTs', name: MyNFTs, component: MyNFTs, meta: { requireAuth: true } },
            { path: '/center/inReview', name: InReview, component: InReview, meta: { requireAuth: true } },
            { path: '/center/mycommittees', name: MyCommittee, component: MyCommittee, meta: { requireAuth: true } },
            { path: '/center/selfCommittee', name: SelfCommittee, component: SelfCommittee, meta: { requireAuth: true } },
            { path: '/center/nftSelling', name: NFTSelling, component: NFTSelling, meta: { requireAuth: true } },
            //{ path: '/center/mint', name: Mint, component: Mint },
            { path: '/center/createcommittees',name:CreateCommittee,component:CreateCommittee,meta:{requireAuth:true}},
            { path: '/center/information', component: Information, meta: { requireAuth: true } },// 首页
            { path: '/center/selfConference', component: SelfConference, meta: { requireAuth: true } },// 
            { path: '/center/inReviewPapers', component: InReviewPapers, meta: { requireAuth: true } },// 
            { path: '/center/ReviewedPapers', component: ReviewedPapers, meta: { requireAuth: true } },// 
            { path: '/center/Reviewed', component: Reviewed, meta: { requireAuth: true } },// 
             { path: '/center/myconference', name: 'MyConference', component: MyConference, meta: { requireAuth: true } }, // 商品管理
             { path: '/center/createConference', name: 'createConference', component: CreateConference, meta: { requireAuth: true } }, // 页面1
            // { path: 'Page2', name: 'page2', component: PageTwo } // 页面2
        ]
    }
]

// 3. 创建 router 实例，然后传 `routes` 配置
// 你还可以传别的配置参数, 不过先这么简单着吧。
const router = new VueRouter({
    mode: 'history',
    routes // (缩写) 相当于 routes: routes
})

// 将 router 实例暴露出去
export default router
