<template>
    <div class="list-wrap">
        <section class="LiveMints_container__0ES6y">
            <div class="LiveMints_section_header_container__hrIL1">
                <div class="LiveMints_header_left">
                    <span class="LiveMints_header_left_txt">Pending List</span>
                </div>
                <div class="LiveMints_header_left">
                    <span class="pause_txt" v-show="mouseOnUl">已暂停</span>
                </div>

                <div class="LiveMints_header_right_container__j6t4N">
                    <button class="LiveMints_btn_pin" aria-expanded="false" aria-haspopup="dialog">
                        <div class="LiveMints_pinned_div">
                            <div class="LiveMints_pinned_pulse____tVQ"></div>
                            <svg-icon icon-class="pin" className="pin-icon"></svg-icon>
                        </div>
                    </button>
                </div>
            </div>
            <div class="LiveMintsFilter_filter_container__n_iEu">
                <div class="LiveMintsFilter_row_container__Z4QDT">
                    <div class="LiveMintsFilter_row__GTDog">
                        <div class="LiveMintsFilter_price_input_container__lFoEF">
                            <span class="LiveMintsFilter_filter_label__GrE2a">Min</span>
                            <input class="LiveMintsFilter_price_range_input__M_xOn" inputmode="decimal" min="0" name="minPrice" placeholder="0" step="0.01" type="number" value="0"></div>
                        <div class="LiveMintsFilter_price_input_container__lFoEF">
                            <span class="LiveMintsFilter_filter_label__GrE2a">Max</span>
                            <input class="LiveMintsFilter_price_range_input__M_xOn" inputmode="decimal" min="0" name="maxPrice" placeholder="100" step="0.01" type="number" value="100"></div>
                            <button class="filter_button LiveMintsFilter_filter_active__bzzUL">
                            <span>Free</span>
                        </button>
                        <button class="filter_button" disabled=""><span>Only</span>
                            <div class="LiveMintsFilter_filter_icon__sHfad">
                                <svg-icon icon-class="verified" className="verify-only-icon"></svg-icon>
                            </div>
                        </button>
                        <button class="filter_button LiveMintsFilter_filter_active__bzzUL">
                            <span>Exclude</span>
                            <div class="LiveMintsFilter_filter_icon__sHfad">
                                <svg-icon icon-class="airdrop" className="verify-only-icon"></svg-icon>
                            </div>
                        </button>
                    </div>
                </div>
            </div>
            <div class="list_container">
                <ul :class="{ list_ul: animate == true }">
                    <li class="list_li"
                      style="transform-origin: 0px 0px;" @mouseenter="showMint(item)" @mouseleave="hideMint(item)" v-for="(item, index) in pendList" :key="index">
                        <div style="display: flex;">
                            <div class="Thumbnail_container" style="width: 42px; height: 42px;">
                                <picture style="width: 42px; height: 42px;">
                                    <source :srcset="imgUrl">
                                    <img :alt="item.nftName" 
                                    base="@/assets/img/no-img.png" size="42" 
                                    :url="item.imgUrl" 
                                    class="Thumbnail_img" 
                                    :src="item.imgUrl" 
                                    width="42" height="42" style="width: 42px; height: 42px;">
                                </picture>
                            </div>
                        </div>
                        <div class="LiveMintsItem_contract_info_container__i23cB">
                            <div class="LiveMintsItem_contract_name_container__QgQAd">
                                <span class="LiveMintsItem_contract_name__RUN2_">{{item.nftName}}</span>
                                <div class="LiveMintsItem_quantity_container__EDehQ">
                                    <span class="LiveMintsItem_quantity__me9ys">{{"x "+item.number}}</span>
                                </div>
                            </div>
                            <!-- //添加按钮 -->
                            <div class="LiveMintsItem_contract_extra_container__aKSiQ">
                                <span>{{"Ξ "+item.value}}</span>
                                <div class="LiveMintsItem_extra_divider__dMi2j"></div>
                                <span>{{"🔥 "+item.gas+" Gwei"}}</span>
                                <div class="LiveMintsItem_extra_divider__dMi2j"></div>
                                <div class="LiveMintsItem_verified_container__sHpm8">
                                    <svg-icon icon-class="verified" className="verify-item"></svg-icon>
                                </div>
                                <span class="LiveMintsItem_function_name__QrhxJ">{{item.mintFunc}}</span>
                            </div>
                            <!-- 添加按钮 -->
                        </div>
                        <div class="LiveMintsItem_buttons_container__eQfZr" v-show="item.isShowMint">
                            <button class="mint_button" @click="clickMint(item.mintHash)">Mint</button>
                        </div>
                        <div class="LiveMintsItem_buttons_container__eQfZr">
                            <a aria-label="live_mints_tx_hash_link" class="ExternalLink_external_link__5BhMS LiveMintsItem_button__opvYi" 
                            href="https://etherscan.io/tx/0xea923e0e49705b4e6d30e8e0023ad098afd19337ba81d59d78e94aee7109d3d9" role="link"
                             rel="noopener noreferrer nofollow" target="_blank">
                                <svg-icon icon-class="etherscan-logo-circle" className="etherscan-icon" ></svg-icon>
                            </a>
                        </div>
                    </li>
                </ul>
            </div>
        </section>
    </div>
</template>


<script>
    export default{
       name:"PendingList"
    }
</script>
<script setup>
import {onMounted,onUpdated,onActivated,onDeactivated,ref} from "vue";
import { emitter } from '@/common/bus'
const pendList = ref([])

const animate= ref(false)
const mouseOnUl = ref(false)
const showMint = (item)=>{
    item.isShowMint= true
    mouseOnUl.value= true
}
const hideMint = (item)=>{
    item.isShowMint= false
    mouseOnUl.value= false
}
const newItem = ref([])
const scroll=()=> {
    if(!mouseOnUl.value){
        animate.value = true; // 为true时将动画样式赋值给布局，见布局处引用
        setTimeout(() => {
          //  这里直接使用了es6箭头写法，省去了处理this指向偏移问题
          pendList.value.unshift(newItem.value[0])
          newItem.value.shift()
          if(pendList.value.length>20){
              pendList.value.pop(); //删除数组的第一个元素
          }
          animate.value = false;
        }, 500);//每条滚动时长0.5s
    }

    }
const getScrollMessage=()=>{
    setInterval(()=>{
        newItem.value.push({
            nftName:"KURENAI -HOOZUKI - "+(Math.random()*1000).toString().slice(0,3),
            imgUrl:"https://i.seadn.io/gcs/files/8347173f9e5bce705308c8705e5f42aa.jpg?w=500&amp;auto=format",
            number:"2",
            value:"1.2",
            gas:"14",
            mintFunc:"wlMint",
            address:"",
            mintHash:"0x24b12f402332724b8b20f18597c511cddc8ce4175e7288777b9a2411d34a7938"
        })
    },2000)
    //这里用网络请求替换真实数据，把下面的代码放到响应成功的线程里
    setInterval(scroll, 2000)
}
const clickMint = (hash)=>{
    emitter.emit('showHashMintPop',(hash))
}
onMounted(() => {
    getScrollMessage()
    pendList.value=[{
    nftName:"KURENAI -HOOZUKI",
    imgUrl:"https://i.seadn.io/gcs/files/8347173f9e5bce705308c8705e5f42aa.jpg?w=500&amp;auto=format",
    number:"2",
    value:"1.2",
    gas:"14",
    mintFunc:"wlMint",
    address:"",
    mintHash:"0x24b12f402332724b8b20f18597c511cddc8ce4175e7288777b9a2411d34a7938"

},
]
})

</script>

<style lang="scss" scoped>
.list-wrap{
    width: 99.9%;
    height: auto;
    min-height: 880px;
    box-shadow: var(--el-box-shadow-light);
    background-color: #fff;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

.list_container {
    width: 100%;
    height: 100%;
    // background-color: var(--gd-primary50);
    overflow-y: scroll;
}
.LiveMints_container__0ES6y {
    display: flex;
    flex-direction: column;
    // width: 433px;
    width: 100%;
    height: 100%;
    min-width: 433px;
    background-color: var(--bg-color);
}
.verify-only-icon{
    height:1em;
    width:1em;
}
.verify-item{
    height:0.9em;
    width:0.9em;
}
.etherscan-icon{
    height:1.5em;
    width:1.5em;
    
}
.etherscan-icon:hover{
    border: 2px #126b40 solid;
    border-radius: 50%;
    background-color: #126b40;
}
.pin-icon{
    height:2em;
    width:2em;
}
.LiveMints_section_header_container__hrIL1 {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 24px 45px 18px;
    // width: 90%;
    height: 78px;
    background-color: var(--bg-color);
}
.LiveMintsFilter_price_range_input__M_xOn {
    padding: 4px 4px 4px 8px;
    width: 64px;
    height: 18px;
    outline: none;
    border: 1px solid var(--gd-gray100) !important;
    border-radius: 4px;
    background-color: var(--bg-color);
    color: var(--gd-gray700);
}
.LiveMints_header_left {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
}
.LiveMints_header_left_txt {
    font-weight: 700;
    font-size: 24px;
    line-height: 30px;
    letter-spacing: -.24px;
    // color: var(--el-text-color-primary);
}
.pause_txt {
    font-size: 14px;
    line-height: 20px;
    letter-spacing: -.24px;
    color: red;
}

.LiveMints_btn_pin, .LiveMints_header_right_container__j6t4N {
    display: flex;
    align-items: center;
    justify-content: center;
}

.Thumbnail_img {
    aspect-ratio: 1;
}
.Thumbnail_img, .Thumbnail_video__wXLg_ {
    border-radius: 50%;
    pointer-events: none;
    -o-object-fit: cover;
    object-fit: cover;
}
.LiveMintsItem_contract_info_container__i23cB {
    display: flex;
    flex-grow: 1;
    flex-direction: column;
    align-items: flex-start;
    margin-left: 12px;
    min-width: 0;
}

.LiveMints_pinned_div {
    position: relative;
    width: 24px;
    height: 24px;
}
.LiveMints_pinned_pulse____tVQ {
    position: absolute;
    right: 4.5px;
    top: 4.5px;
    display: grid;
    width: 16px;
    height: 16px;
    border-radius: 50%;
    transition: all .2s;
    animation: LiveMints_pinned_pulse____tVQ 1.2s ease-out 0s infinite;
}



.LiveMintsItem_function_name__QrhxJ {
    color: var(--gd-primary500);
    text-overflow: ellipsis;
    overflow: hidden;
}
.LiveMints_header_right_container__j6t4N {
    flex-direction: row;
}
.LiveMintsFilter_filter_container__n_iEu {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 0 45px 12px;
}
.LiveMintsFilter_row__GTDog, .LiveMintsFilter_row_container__Z4QDT {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    width: 100%;
}
.LiveMintsItem_contract_extra_container__aKSiQ {
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 100%;
    white-space: nowrap;
}
.LiveMintsFilter_row__GTDog, .LiveMintsFilter_row_container__Z4QDT {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    width: 100%;
}
.LiveMintsFilter_price_input_container__lFoEF {
    display: flex;
    flex-direction: row;
    align-items: center;
}
.LiveMintsItem_quantity__me9ys {
    padding: 0 4px;
    font-weight: 500;
    font-size: 10px;
    line-height: 18px;
    letter-spacing: -.24px;
    color: var(--bg-color);
}
.LiveMintsFilter_filter_label__GrE2a {
    padding-right: 4px;
    font-weight: 500;
    font-size: 12px;
    line-height: 18px;
    letter-spacing: -.24px;
    color: var(--gd-gray850);
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
}
.LiveMintsFilter_price_input_container__lFoEF {
    display: flex;
    flex-direction: row;
    align-items: center;
}
.LiveMintsItem_quantity_container__EDehQ {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0 4px;
    margin-left: 6px;
    height: 16px;
    background-color: var(--gd-primary500);
    border-radius: 14px;
}
.LiveMintsFilter_filter_active__bzzUL {
    border-color: var(--gd-primary500);
    background-color: var(--gd-primary50);
    color: var(--gd-primary500);
}
.filter_button {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4px 8px;
    height: 28px;
    border: 1px solid var(--gd-gray100);
    border-radius: 4px;
    background-color: var(--bg-color);
    font-size: 12px;
    line-height: 18px;
    letter-spacing: -.24px;
    color: var(--gd-gray850);
    text-align: center;
}
.mint_button {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4px 8px;
    height: 22px;
    width: 45px;
    border: 1px solid var(--gd-primary500);
    border-radius: 6px;
    background-color: var(--gd-primary500);
    font-size: 14px;
    line-height: 24px;
    margin-right: 100%;
    color: var(--gd-gray850);
    text-align: center;
    cursor: pointer;
}
.mint_button:hover{
    color: var(--bg-color);
    background-color: var(--gd-primary500);

}
.LiveMintsFilter_filter_icon__sHfad {
    margin-left: 4px;
    margin-top: 3px;
}
.list_ul {
    transition: all 0.5s;
  /*几个高度要一致*/
  margin-top: 74px;
//   padding-top: 74px;
    display: flex;
    flex-direction: column-reverse;
}

ul {
    list-style: none;
}
.list_li {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 16px 45px;
    height: 42px;
    // border-bottom: 2px #d3d4c4 solid;
    transition:color .5s ease-out, background .5s ease-in-out;
    box-sizing: content-box;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    cursor: pointer !important;
}
.list_li:hover{
    background-color: var(--gd-gray100);
}
.Thumbnail_container {
    border-radius: 50%;
    background-color: var(--gd-gray50);
    line-height: 1;
}
.LiveMintsItem_contract_name_container__QgQAd {
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-bottom: 2px;
    width: 100%;
}
.LiveMintsItem_contract_extra_container__aKSiQ, .LiveMintsItem_time__TU2wG {
    font-size: 12px;
    line-height: 18px;
    letter-spacing: -.24px;
    color: var(--gd-gray700);
}
.LiveMintsItem_contract_name__RUN2_ {
    font-weight: 500;
    font-size: 14px;
    line-height: 22px;
    letter-spacing: -.24px;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
}
.LiveMintsItem_extra_divider__dMi2j {
    margin: 0 6px;
    width: 1px;
    height: 8px;
    background-color: var(--gd-gray700);
}
.LiveMintsItem_extra_divider__dMi2j {
    margin: 0 6px;
    width: 1px;
    height: 8px;
    background-color: var(--gd-gray700);
}
.LiveMintsItem_verified_container__sHpm8 {
    padding-right: 4px;
}
.LiveMintsItem_airdrop_container__VDag3, .LiveMintsItem_verified_container__sHpm8 {
    display: flex;
    align-items: center;
    justify-content: center;
}
.LiveMintsItem_buttons_container__eQfZr {
    display: flex;
    flex-direction: row;
    align-items: center;
}
.LiveMintsItem_button__opvYi {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-left: 8px;
    width: 28px;
    height: 28px;
    border-radius: 50%;
    background-color: var(--gd-gray50);
    opacity: 1;
    cursor: pointer;
}

.ExternalLink_external_link__5BhMS {
    display: inline-flex;
    align-items: center;
    text-decoration: underline;
}
a {
    all: unset;
}
blockquote, body, dd, dl, dt, fieldset, figure, h1, h2, h3, h4, h5, h6, hr, html, iframe, legend, li, ol, p, pre, textarea, ul {
    margin: 0;
    padding: 0;
}
// * {
//     -webkit-tap-highlight-color: rgba(255,255,255,0);
//     -webkit-touch-callout: none;
// }
// *, :after, :before {
//     box-sizing: inherit;
// }
user agent stylesheet
li {
    display: list-item;
    text-align: -webkit-match-parent;
}
ul {
    list-style: none;
}
user agent stylesheet
ul {
    list-style-type: disc;
}
</style>