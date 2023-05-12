<template lang="pug">
.container
    Header
    .main
        .sidebar
        router-view
</template>

<script lang="ts" setup>
import Header from '@/views/layout/Header.vue';
</script>

<style lang="scss" scoped>
.container {
    width: 100vw;
    height: 100vh;
}
</style>