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
import SelfCommittee from '../views/center/SelfCommittee.vue'
import SelfConference from '../views/center/SelfConference.vue'
import InReviewPapers from '@/views/center/InReviewPapers.vue'
import InReview from '@/views/center/InReview.vue'
import Reviewed from '@/views/center/Reviewed.vue'
import ReviewedPapers from '@/views/center/ReviewedPapers.vue'
import Mint from '../views/center/Mint.vue'
import Papers from '../views/center/Papers.vue'
import Reviews from '../views/center/Reviews.vue'
import Users from '../views/center/Users.vue'
import CreateJournal from '../views/center/CreateJournal.vue'
import SelfJournal from '../views/center/SelfJournal.vue'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        redirect: "/home",
        name: 'main',
        component: Main,
        children: [
            { path: 'home', name: 'home', component: Home },
            { path: 'committees', name: 'committees', component: Committees },
            { path: 'detailCommittee', name: 'homeDetailCommittee', component: HomeDetailCommittee },
            { path: 'conferences', name: 'conferences', component: Conferences },
            { path: 'detailConference', name: 'homeDetailConference', component: HomeDetailConference },
            { path: 'journals', name: 'journals', component: Journals },
            { path: 'market', name: 'market', component: Market },
            { path: 'detailJournal', name: 'homeDetailJournal', component: HomeDetailJournal },
            { path: 'conferenceSubmit', name: 'conferenceSubmit', component: ConferenceSubmit },
            { path: 'journalSubmit', name: 'journalSubmit', component: JournalSubmit },
            { path: 'journalIssues', name: 'journalIssues', component: JournalIssues },
            { path: 'conferenceIssues', name: 'conferenceIssues', component: ConferenceIssues },
            { path: 'publications', name: 'publications', component: Publications },
            { path: 'detailPaper', name: 'homeDetailPaper', component: HomeDetailPaper },
            { path: 'user', name: 'user', component: User },
            { path: 'test', name: 'test', component: Test }
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
            { path: '/center/detailCommittee', name: 'detailCommittee', component: DetailCommittee, meta: { requireAuth: true } },
            { path: '/center/detailConference', name: 'detailConference', component: DetailConference, meta: { requireAuth: true } },
            { path: '/center/detailJournal', name: 'detailJournal', component: DetailJournal, meta: { requireAuth: true } },
            { path: '/center/updateInformation', name: 'updateInformation', component: UpdateInformation, meta: { requireAuth: true } },
            { path: '/center/updateCommittee', name: 'updateCommittee', component: UpdateCommittee, meta: { requireAuth: true } },
            { path: '/center/updateJournal', name: 'updateJournal', component: UpdateJournal, meta: { requireAuth: true } },
            { path: '/center/updateConference', name: 'updateConference', component: UpdateConference, meta: { requireAuth: true } },
            { path: '/center/detailPaper', name: 'detailPaper', component: DetailPaper, meta: { requireAuth: true } },
            { path: '/center/updatePaper', name: 'updatePaper', component: UpdatePaper, meta: { requireAuth: true } },
            { path: '/center/revisePaper', name: 'revisePaper', component: RevisePaper, meta: { requireAuth: true } },
            { path: '/center/reviewPaperDetail', name: 'reviewPaperDetail', component: ReviewPaperDetail, meta: { requireAuth: true } },
            { path: '/center/journalIssues', name: 'journalIssuesCenter', component: JournalIssuesCenter, meta: { requireAuth: true } },
            { path: '/center/conferenceIssues', name: 'conferenceIssuesCenter', component: ConferenceIssuesCenter, meta: { requireAuth: true } },
            { path: '/center/myNFTs', name: 'myNFTs', component: MyNFTs, meta: { requireAuth: true } },
            { path: '/center/inReview', name: 'inReview', component: InReview, meta: { requireAuth: true } },
            { path: '/center/mycommittees', name: 'myCommittee', component: MyCommittee, meta: { requireAuth: true } },
            { path: '/center/selfCommittee', name: 'selfCommittee', component: SelfCommittee, meta: { requireAuth: true } },
            { path: '/center/nftSelling', name: 'nftSelling', component: NFTSelling, meta: { requireAuth: true } },
            { path: '/center/createcommittees', name: 'createCommittee', component: CreateCommittee, meta: { requireAuth: true } },
            { path: '/center/information', name: 'information', component: Information, meta: { requireAuth: true } },
            { path: '/center/selfConference', name: 'selfConference', component: SelfConference, meta: { requireAuth: true } },
            { path: '/center/inReviewPapers', name: 'inReviewPapers', component: InReviewPapers, meta: { requireAuth: true } },
            { path: '/center/ReviewedPapers', name: 'reviewedPapers', component: ReviewedPapers, meta: { requireAuth: true } },
            { path: '/center/Reviewed', name: 'reviewed', component: Reviewed, meta: { requireAuth: true } },
            { path: '/center/myconference', name: 'myConference', component: MyConference, meta: { requireAuth: true } },
            { path: '/center/createConference', name: 'createConference', component: CreateConference, meta: { requireAuth: true } },
            { path: '/center/papers', name: 'papers', component: Papers, meta: { requireAuth: true } },
            { path: '/center/reviews', name: 'reviews', component: Reviews, meta: { requireAuth: true } },
            { path: '/center/users', name: 'users', component: Users, meta: { requireAuth: true } },
            { path: '/center/mint', name: 'mint', component: Mint, meta: { requireAuth: true } },
            { path: '/center/createJournal', name: 'createJournal', component: CreateJournal },
            { path: '/center/selfJournal', name: 'selfJournal', component: SelfJournal }
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
