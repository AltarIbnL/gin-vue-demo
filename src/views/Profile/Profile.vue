<template>
  <h3> 您好 <span v-if="userinfo">{{userinfo.name}}</span>></h3>
</template>
<script>
import {mapState} from "vuex";

export default {
  computed: mapState({
    // 从缓存中获取改为从vuex 中获取，因为vuex是响应式的，
    userinfo: (state) => state.userModule.userInfo,
    // return this.$store.state.userModule.userInfo;
  }),
};

</script>
<style lang="scss" scoped></style>
