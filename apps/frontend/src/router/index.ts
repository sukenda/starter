import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/features/auth/sessionStore'

export const router=createRouter({history:createWebHistory(import.meta.env.BASE_URL),routes:[
 {path:'/login',name:'login',component:()=>import('@/features/auth/pages/LoginPage.vue'),meta:{public:true}},
 {path:'/',component:()=>import('@/layouts/AdminLayout.vue'),children:[
  {path:'',name:'dashboard',component:()=>import('@/features/dashboard/pages/DashboardPage.vue')},
  {path:'users',name:'users',component:()=>import('@/features/admin/pages/UsersPage.vue'),meta:{permission:'users.read'}},
  {path:'roles',name:'roles',component:()=>import('@/features/admin/pages/RolesPage.vue'),meta:{permission:'roles.read'}},
  {path:'menus',name:'menus',component:()=>import('@/features/admin/pages/MenusPage.vue'),meta:{permission:'menus.read'}},
 ]},
]})
router.beforeEach(async to=>{const auth=useAuthStore();if(to.meta.public){if(auth.isAuthenticated&&to.name==='login')return '/';return true}if(!auth.isAuthenticated)return {name:'login',query:{redirect:to.fullPath}};const permission=to.meta.permission as string|undefined;if(permission&&!auth.hasPermission(permission))return '/';return true})
