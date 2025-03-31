<template>
  <div class="error-page">
    <div class="error-content">
      <div class="error-image">
        <img src="@/assets/images/403.svg" alt="403" />
      </div>
      <h1 class="error-title">403</h1>
      <div class="error-desc">抱歉，您没有访问该页面的权限</div>
      <div class="error-actions">
        <a-button type="primary" @click="$router.push('/')">
          返回首页
        </a-button>
        <a-button style="margin-left: 16px" @click="$router.go(-1)">
          返回上一页
        </a-button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.error-page {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-color: var(--color-bg-1);
}

.error-content {
  text-align: center;
  padding: 20px;
}

.error-image {
  max-width: 400px;
  margin: 0 auto 32px;
}

.error-image img {
  width: 100%;
  height: auto;
}

.error-title {
  font-size: 72px;
  font-weight: 600;
  color: var(--color-text-1);
  margin: 0 0 16px;
}

.error-desc {
  font-size: 20px;
  color: var(--color-text-3);
  margin-bottom: 24px;
}

.error-actions {
  margin-top: 32px;
}
</style> 