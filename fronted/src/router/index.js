import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

/* Layout */
import Layout from "@/layout";
import Blayout from "@/blayout";

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: "/login",
    component: () => import("@/views/admin/login/index"),
    hidden: true
  },

  {
    path: "/404",
    component: () => import("@/views/404"),
    hidden: true
  },
  {
    path: "/",
    component: Blayout,
    children: [
      {
        path: "",
        name: "Home",
        hidden: true,
        component: () => import("@/views/blog/home"),
        meta: { title: "主页" }
      }
    ]
  },
  {
    path: "/articles",
    component: Blayout,
    children: [
      {
        path: "",
        name: "CTQArticle",
        hidden: true,
        component: () => import("@/views/blog/article/article"),
        meta: { title: "文章列表" }
      }
    ]
  },
  {
    path: "/category",
    component: Blayout,
    children: [
      {
        path: "",
        name: "Bcategory",
        hidden: true,
        component: () => import("@/views/blog/category"),
        meta: { title: "分类" }
      }
    ]
  },
  {
    path: "/tag",
    component: Blayout,
    children: [
      {
        path: "",
        name: "Btag",
        hidden: true,
        component: () => import("@/views/blog/tag"),
        meta: { title: "标签" }
      }
    ]
  },
  {
    path: "/about",
    component: Blayout,
    children: [
      {
        path: "",
        name: "About",
        hidden: true,
        component: () => import("@/views/blog/about"),
        meta: { title: "关于" }
      }
    ]
  },
  {
    path: "/admin",
    component: Layout,
    redirect: "/admin/dashboard",
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        component: () => import("@/views/admin/dashboard/index"),
        meta: { title: "概览", icon: "dashboard" }
      }
    ]
  },
  {
    path: "/admin/article/create",
    component: Layout,
    children: [
      {
        path: "",
        name: "CreateArticle",
        component: () => import("@/views/admin/article/create"),
        meta: { title: "添加文章", icon: "form" }
      }
    ]
  },

  {
    path: "/admin/article",
    component: Layout,
    children: [
      {
        path: "list",
        name: "Article",
        component: () => import("@/views/admin/article/list"),
        meta: { title: "文章列表", icon: "list" }
      },
      {
        path: "edit/:id(\\d+)",
        name: "EditArticle",
        hidden: true,
        component: () => import("@/views/admin/article/edit"),
        meta: { title: "编辑文章" }
      }
    ]
  },

  {
    path: "/admin/category",
    component: Layout,
    children: [
      {
        path: "list",
        name: "Category",
        component: () => import("@/views/admin/category/index"),
        meta: { title: "分类管理", icon: "component" }
      }
    ]
  },
  {
    path: "/admin/tag",
    component: Layout,
    children: [
      {
        path: "list",
        name: "Tag",
        component: () => import("@/views/admin/tag/index"),
        meta: { title: "标签管理", icon: "tag" }
      }
    ]
  },
  {
    path: "/admin/soup",
    component: Layout,
    children: [
      {
        path: "list",
        name: "Soup",
        component: () => import("@/views/admin/soup/index"),
        meta: { title: "鸡汤管理", icon: "table" }
      }
    ]
  },
  {
    path: "/admin/comment",
    component: Layout,
    children: [
      {
        path: "list",
        name: "Comment",
        component: () => import("@/views/admin/comment/index"),
        meta: { title: "评论管理", icon: "comment" }
      }
    ]
  },
  {
    path: "/admin/about",
    component: Layout,
    children: [
      {
        path: "",
        name: "Aabout",
        component: () => import("@/views/admin/about/index"),
        meta: { title: "编辑关于页", icon: "about" }
      }
    ]
  },

  {
    path: "/admin/profile",
    component: Layout,
    children: [
      {
        path: "",
        name: "Profile",
        component: () => import("@/views/admin/profile/index"),
        meta: { title: "个人中心", icon: "user" }
      }
    ]
  },

  // 404 page must be placed at the end !!!
  { path: "*", redirect: "/404", hidden: true }
];

const createRouter = () =>
  new Router({
    // mode: 'history', // require service support
    scrollBehavior: () => ({ y: 0 }),
    routes: constantRoutes
  });

const router = createRouter();

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter();
  router.matcher = newRouter.matcher; // reset router
}

export default router;
