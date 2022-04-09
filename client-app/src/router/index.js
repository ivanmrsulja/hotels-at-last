import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const roles = { admin: 1, user: 0 };

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/HomeView.vue"),
    meta: {
      authenticated: false,
      authorities: [],
    },
  },
  {
    path: "/hotels/:id",
    name: "DetailedHotel",
    component: () => import("../views/DetailedHotelView.vue"),
    meta: {
      authenticated: false,
      authorities: [],
    },
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/Login.vue"),
    meta: {
      authenticated: false,
      authorities: [],
    },
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../views/Register.vue"),
    meta: {
      authenticated: false,
      authorities: [],
    },
  },
  {
    path: "/reported-comments",
    name: "ReportedComments",
    component: () => import("../views/ReportedCommentsView.vue"),
    meta: {
      authenticated: true,
      authorities: [roles.admin],
    },
  },
  {
    path: "/hotels-admin",
    name: "HotelsAdministration",
    component: () => import("../views/HotelsAdminView.vue"),
    meta: {
      authenticated: true,
      authorities: [roles.admin],
    },
  },
  {
    path: "/reservations",
    name: "ReservationList",
    component: () => import("../views/ReservationListView.vue"),
    meta: {
      authenticated: true,
      authorities: [roles.user],
    },
  },
];

const router = new VueRouter({
  routes,
});

export default router;
