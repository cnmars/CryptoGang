<template>
    <div class="group">
      <svg class="icon" aria-hidden="true" viewBox="0 0 24 24"><g><path d="M21.53 20.47l-3.66-3.66C19.195 15.24 20 13.214 20 11c0-4.97-4.03-9-9-9s-9 4.03-9 9 4.03 9 9 9c2.215 0 4.24-.804 5.808-2.13l3.66 3.66c.147.146.34.22.53.22s.385-.073.53-.22c.295-.293.295-.767.002-1.06zM3.5 11c0-4.135 3.365-7.5 7.5-7.5s7.5 3.365 7.5 7.5-3.365 7.5-7.5 7.5-7.5-3.365-7.5-7.5z"></path></g></svg>
      <input :placeholder="$t('navbar.ph.search')" type="search" class="input">
    </div>
</template>
<script>
    export default{
        name:"SearchBar"
    }
</script>
<script setup>
    import { useI18n } from 'vue-i18n'
    const { t } = useI18n()

</script>
    
    <style lang="scss" scoped>
    .group {
     display: flex;
     line-height: 28px;
     align-items: center;
     position: relative;
     max-width: 990px;
    }
    
    .input {
     width: 100%;
     height: 34px;
     line-height: 28px;
     padding: 0 1rem;
     padding-left: 2.5rem;
     border: 1px solid transparent;
     border-radius: 6px;
     outline: none;
     background-color: var(--gd-search-input-bg);
     color: var(--el-input-text-color);
     transition: .3s ease;
    }
    
    .input::placeholder {
     color: #9e9ea7;
    }
    
    .input:focus, input:hover {
     outline: none;
     border-color: var(--el-color-primary);
     background-color: var(--gd-search-input-hover);
    //  box-shadow: 0 0 0 4px rgba(12, 137, 60, 0.1);
    }
    
    .icon {
     position: absolute;
     left: 1rem;
     fill: #9e9ea7;
     width: 1rem;
     height: 1rem;
    }
    
    
    
    
    
    
    </style>