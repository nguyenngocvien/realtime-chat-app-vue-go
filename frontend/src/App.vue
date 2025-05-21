<template>
  <div>
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import TopBar from './components/TopBar.vue'

export default defineComponent({
  name: 'App',
  components: { TopBar },
  data() {
    return {
      isSidebarOpen: false
    }
  }
})
</script>

<style scoped>
aside {
  width: 300px;
  top: 64px;
  bottom: 0;
}

@media (min-width: 768px) {
  aside {
    top: 0;
  }
}
</style>