<template>
  <transition name="nothing" v-bind="$attrs" v-on="$listeners">
    <slot></slot>
  </transition>
</template>

<script>
export default {
  props: {
    wrapperId: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      innerHTML: ""
    }
  },
  beforeMount() {
    const elem = document.getElementById(this.wrapperId);
    this.innerHTML = elem.innerHTML;
  },
  mounted() {
    const elem = document.getElementById(this.wrapperId);
    elem.innerHTML = this.innerHTML;
  }
}
</script>