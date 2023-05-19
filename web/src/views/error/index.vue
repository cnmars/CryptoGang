<template>
  <div>
    <div class="big">
      <div class="inner">
        <!-- <img src="../../assets/svg/404.svg"> -->
        <svg-icon icon-class="404" className="svg404"></svg-icon>
        <!-- <p>页面被神秘力量吸走了</p> -->
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Error'

}
</script>

<style lang="scss">
    .big{
        width: 100%;
        height: calc(100vh - 260px);
        background-color: rgb(244,244,244);
        position: relative;
    }
    .inner{
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
    .inner p {
        text-align: center;
        font-size: 24px;
    }
    .inner .leftPic{
        width: 60px;
        height: 60px;
        margin-left: 44%;
        margin-top: 20px;
    }

    .svg404{
      width: 33em;
      height: 33em;
      color: #126b40;
    }
</style>
