<template>
  <div id="app">
      <navbar />
    <b-container>
    <router-view/>
    </b-container>
  </div>
</template>
<script>
// eslint-disable-next-line import/extensions
import Navbar from '@/views/layout/Navbar';

export default {
  components: { Navbar },
  data() {
    return {};
  },
};

</script>
<style lang="scss">

</style>
