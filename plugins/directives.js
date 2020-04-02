import Vue from "vue";

Vue.directive('ignore',
{
  bind(el, binding, vnode) {
      console.log("bind");
      // console.log( el.innerHTML )
  },
  inserted(el, binding, vndoe) {
    console.log("inserted");
  },
 updated(el, binding, vnode) {
     console.log("updated");
 },
 componentUpdated(el, binding, vnode, oldVnode) {
     console.log("componentUpdated");
 }
}
);
