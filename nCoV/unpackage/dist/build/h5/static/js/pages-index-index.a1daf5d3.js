(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-index-index"],{"0a24":function(t,e,i){var a=i("a374");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var n=i("4f06").default;n("2f6f7379",a,!0,{sourceMap:!1,shadowMode:!1})},"502f":function(t,e,i){"use strict";var a=i("0a24"),n=i.n(a);n.a},5764:function(t,e,i){"use strict";(function(t){var a=i("288e");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n=a(i("d3ba")),s={data:function(){return{StatusBar:this.StatusBar,tabList:[{title:"实时数据"}],menuList:[{title:"加载中...",url:"#",image:"http://5b0988e595225.cdn.sohucs.com/images/20170831/ea2692adffd64b82a2d5b8f4d3d8d7d8.gif"}],show:!0,allList:[{item:[]}],desc:{},TabCur:0,scrollLeft:0,CustomBar:this.CustomBar,numData:{confirmed:0,suspected:0,serious:0,death:0,cure:0}}},computed:{style:function(){var t=this.StatusBar,e=this.CustomBar,i=this.bgImage,a="height:".concat(e,"px;padding-top:").concat(t,"px;");return this.bgImage&&(a="".concat(a,"background-image:url(").concat(i,");")),a}},methods:{tabSwiper:function(t){this.TabCur=t.detail.current,this.scrollLeft=60*(t.detail.current-1)},tabSelect:function(t){this.TabCur=t.currentTarget.dataset.id,this.scrollLeft=60*(t.currentTarget.dataset.id-1)},getMore:function(t){var e=this;n.default.re("original/tx/").then(function(t){e.desc=t.data.newslist[0].desc})},onLoad:function(t){this.getNum(),this.getNav(),this.baidu(),this.getMore(),this.getMenu()},onReachBottom:function(t){},onPageScroll:function(t){},onHide:function(){},onShow:function(){},onUnload:function(){},onPullDownRefresh:function(){},onBackPress:function(){},getNav:function(){var t=this;n.default.re("nav/app/").then(function(e){t.allList=n.default.acc(t.allList,e.data.data);var i=e.data.data.map(function(t,e){return{title:t.title,id:e}});t.tabList=n.default.acc(t.tabList,i)})},getMenu:function(){var t=this;n.default.re("nav/recommendation").then(function(e){t.menuList=e.data.data})},getNum:function(){var e=this;n.default.re("latest/tx/").then(function(i){var a=i.data.data.count;a.update=n.default.toTimes(i.data.data.update_time),e.numData=a,e.moreData=i.data.data.info;var s={date:n.default.getYesterdayDate()};n.default.re("daily/sum/",s).then(function(e){t.log(e)})})},baidu:function(){(function(){var t=document.createElement("script");t.src="https://hm.baidu.com/hm.js?cc2f5bf672651df7ddf7d6093ab34350";var e=document.getElementsByTagName("script")[0];e.parentNode.insertBefore(t,e)})()},goJump:function(e){t.log(e),uni.navigateTo({url:e})},tz:function(t){n.default.tzWeb(t)}}};e.default=s}).call(this,i("5a52")["default"])},a374:function(t,e,i){e=t.exports=i("2350")(!1),e.push([t.i,".user-view[data-v-6bb2b3c2]{z-index:-1;width:100%;height:%?260?%}.navigation-list[data-v-6bb2b3c2]{padding-left:%?32?%;background:#fff;margin-top:%?36?%}.navigation-item[data-v-6bb2b3c2]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-pack:justify;-webkit-justify-content:space-between;justify-content:space-between;-webkit-box-align:center;-webkit-align-items:center;align-items:center;padding-right:%?32?%;padding-left:%?32?%;height:%?106?%;border-bottom:%?1?% solid #f1f1f1}.navigation-item .left[data-v-6bb2b3c2]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center}.navigation-item .right[data-v-6bb2b3c2]{-webkit-box-align:right;-webkit-align-items:right;align-items:right}.navigation-text[data-v-6bb2b3c2]{margin-left:%?21?%;font-size:%?30?%;color:#999;font-weight:500;line-height:%?38?%}.icon-enter[data-v-6bb2b3c2]{width:%?14?%;height:%?24?%}.button-null[data-v-6bb2b3c2]{position:relative;padding-left:%?0?%;background:transparent}.button-hover[data-v-6bb2b3c2]{color:rgba(0,0,0,.6);background-color:hsla(0,0%,87.1%,0)}.button-null[data-v-6bb2b3c2]:after{border:none}.mini-img[data-v-6bb2b3c2]{font-size:%?42?%;color:#888}.mini-img2[data-v-6bb2b3c2]{font-size:%?32?%;color:#888}.text-cuts[data-v-6bb2b3c2]{white-space:nowrap;overflow:hidden}",""])},ac30:function(t,e,i){"use strict";i.r(e);var a=i("5764"),n=i.n(a);for(var s in a)"default"!==s&&function(t){i.d(e,t,function(){return a[t]})}(s);e["default"]=n.a},d3ba:function(t,e,i){"use strict";(function(t){var a=i("288e");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n=a(i("e814")),s=a(i("795b")),c="https://ncov-api.werty.cn:2021/";function u(e,i,a){return new s.default(function(n,s){uni.request({url:c+e,data:i||{},method:a||"GET",success:function(t){n(t)},fail:function(e){t.error(e.data),s(e.data)},error:function(t){return s("网络出错")}})})}function r(){var t=h(1,12);uni.showLoading({title:"更新中..."}),setTimeout(function(){uni.hideLoading(),uni.showModal({title:"温馨提醒",content:"您现在已是最新版本！",showCancel:!1,confirmText:"我知道了"})},579*t)}function o(){var t=new Date;t.setTime(t.getTime()-864e5);var e=t.getMonth()+1;e<10&&(e="0"+e);var i=t.getDate();return i<10&&(i="0"+i),t.getFullYear()+"-"+e+"-"+i}function l(){d("清除成功")}function d(t){uni.showToast({title:t||"错误",duration:2e3,icon:"none"})}function v(){return new s.default(function(e,i){uni.getStorageSync("token")?uni.chooseImage({success:function(i){var a=i.tempFilePaths,n=uni.uploadFile({url:c+"common/upload",filePath:a[0],name:"file",formData:{},success:function(t){d("上传成功"),e({img:JSON.parse(t.data).data.url,url:i.tempFilePaths[0]})},fail:function(e){d("上传失败"),t.error(e)}});n.onProgressUpdate(function(e){t.log(e)})}}):(d("您还没有登录"),i({data:"未登录!"}))})}function f(t){t?uni.navigateTo({url:t}):d("参数错误")}function g(t){t?uni.switchTab({url:t}):uni.switchTab({url:"/pages/index/index.vue"})}function m(t,e){if(e)return(t||[]).concat(e||[]);d("没有更多了...")}function x(t){t?uni.navigateBack({delta:t}):uni.navigateBack({delta:1})}function b(t){uni.setNavigationBarTitle({title:t||""})}function w(){setTimeout(function(){x()},1737)}function h(t,e){switch(arguments.length){case 1:return(0,n.default)(Math.random()*t+1,10);case 2:return(0,n.default)(Math.random()*(e-t+1)+t,10);default:return 0}}function p(t,e,i,a){uni.showModal({title:e||"温馨提醒",content:t||"错误",showCancel:i||!1,confirmText:a||"我知道了"})}function _(){return Date.parse(new Date)}function C(t){uni.reLaunch({url:t})}function y(t){if(t)if(t<1546272e6)e=new Date(1e3*t);else e=new Date(t);else var e=new Date;var i=[e.getFullYear(),e.getMonth()+1,e.getDate()],a=i[0],n=i[1],s=i[2];return n=n<10?"0"+n:n,s=s<10?"0"+s:s,a+"-"+n+"-"+s}function k(t){if(t<1546272e6)var e=new Date(1e3*t);else e=new Date(t);var i=[e.getFullYear(),e.getMonth()+1,e.getDate()],a=(i[0],i[1]),n=i[2],s=[e.getHours(),e.getMinutes(),e.getSeconds()],c=s[0],u=s[1],r=s[2];return a=a<10?"0"+a:a,n=n<10?"0"+n:n,c=c<10?"0"+c:c,u=u<10?"0"+u:u,r=r<10?"0"+r:r,a+"月"+n+"日"+c+":"+u}function D(t){window.location.href=t}function T(t){if(t<1546272e6)var e=new Date(1e3*t);else e=new Date(t);var i=[e.getFullYear(),e.getMonth()+1,e.getDate()],a=i[0],n=i[1],s=i[2],c=[e.getHours(),e.getMinutes(),e.getSeconds()],u=c[0],r=c[1],o=c[2];return n=n<10?"0"+n:n,s=s<10?"0"+s:s,u=u<10?"0"+u:u,r=r<10?"0"+r:r,o=o<10?"0"+o:o,a+"-"+n+"-"+s+" "+u+":"+r+":"+o}var L={updataBtn:r,toTimes:T,upDataImage:v,getTime:_,tzWeb:D,getYesterdayDate:o,toTime:y,acc:m,tz:f,tips:p,rel:C,sw:g,setNavTitle:b,clearCache:l,toast:d,re:u,back:x,toMiniTime:k,sellpBcak:w,randomNum:h};e.default=L}).call(this,i("5a52")["default"])},e588:function(t,e,i){"use strict";i.r(e);var a=i("f6e8"),n=i("ac30");for(var s in n)"default"!==s&&function(t){i.d(e,t,function(){return n[t]})}(s);i("502f");var c=i("f0c5"),u=Object(c["a"])(n["default"],a["a"],a["b"],!1,null,"6bb2b3c2",null);e["default"]=u.exports},f6e8:function(t,e,i){"use strict";var a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",{staticClass:"page"},[i("v-uni-scroll-view",{staticClass:"bg-white fixed nav margin-top-xs",class:t.show?"":"fixed",attrs:{"scroll-x":!0,"scroll-with-animation":!0,"scroll-left":t.scrollLeft}},t._l(t.tabList,function(e,a){return i("v-uni-view",{key:a,staticClass:"cu-item",class:a==t.TabCur?"text-orange cur":"",attrs:{"data-id":a,"data-type":e.id},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.tabSelect.apply(void 0,arguments)}}},[t._v(t._s(e.title))])}),1),i("v-uni-swiper",{staticStyle:{height:"100vh",width:"100vw","padding-top":"100upx"},attrs:{"indicator-dots":!1,autoplay:!1,interval:1e3,current:t.TabCur},on:{change:function(e){arguments[0]=e=t.$handleEvent(e),t.tabSwiper.apply(void 0,arguments)}}},t._l(t.allList,function(e,a){return i("v-uni-swiper-item",{key:a},[i("v-uni-scroll-view",{staticStyle:{height:"100vh",width:"100vw"},attrs:{"scroll-y":"true"}},[0==a?i("v-uni-view",[i("v-uni-view",{staticClass:"cu-card padding-0"},[i("v-uni-view",{staticClass:"lee-titles solid-bottom text-content padding-bottom-sm"},[t._v("常用导航")]),i("v-uni-view",{staticClass:"grid col-5 margin-bottom-sm"},t._l(t.menuList,function(e,a){return i("v-uni-view",{key:a,staticClass:"padding-xs text-center",on:{click:function(i){arguments[0]=i=t.$handleEvent(i),t.tz(e.url)}}},[i("v-uni-image",{staticClass:"cu-avatar lg round",attrs:{src:e.image}}),i("v-uni-view",{staticClass:"text-gray text-sm margin-top-xs text-cuts"},[t._v(t._s(e.title))])],1)}),1)],1),i("v-uni-view",{staticClass:"cu-card padding-0"},[i("v-uni-view",{staticClass:"grid col-5"},[i("v-uni-view",{staticClass:"padding-tb-sm solid-right"},[i("v-uni-view",{staticClass:"text-bold text-red text-center text-xl "},[t._v(t._s(t.numData.confirmed))]),i("v-uni-view",{staticClass:"text-sm text-center margin-top-xs"},[t._v("确诊")]),i("v-uni-view",{staticClass:"text-grey text-sm text-center padding-top-sm"},[t._v("较昨日")]),i("v-uni-view",{staticClass:"margin-top-xs text-red text-center text-xs text-bold"},[t._v(t._s(t.numData.confirmed_incr>0?"+":"")+t._s(t.numData.confirmed_incr))])],1),i("v-uni-view",{staticClass:"padding-tb-sm solid-right"},[i("v-uni-view",{staticClass:"text-bold text-orange text-center text-xl "},[t._v(t._s(t.numData.suspected))]),i("v-uni-view",{staticClass:"text-sm text-center margin-top-xs"},[t._v("疑似")]),i("v-uni-view",{staticClass:"text-grey text-sm text-center padding-top-sm"},[t._v("较昨日")]),i("v-uni-view",{staticClass:"margin-top-xs text-orange text-center text-xs text-bold"},[t._v(t._s(t.numData.suspected_incr>0?"+":"")+t._s(t.numData.suspected_incr))])],1),i("v-uni-view",{staticClass:"padding-tb-sm solid-right"},[i("v-uni-view",{staticClass:"text-bold text-brown text-center text-xl "},[t._v(t._s(t.numData.serious))]),i("v-uni-view",{staticClass:"text-sm text-center margin-top-xs"},[t._v("重症")]),i("v-uni-view",{staticClass:"text-grey text-sm text-center padding-top-sm"},[t._v("较昨日")]),i("v-uni-view",{staticClass:"margin-top-xs text-brown text-center text-xs text-bold"},[t._v(t._s(t.numData.serious_incr>0?"+":"")+t._s(t.numData.serious_incr))])],1),i("v-uni-view",{staticClass:"padding-tb-sm solid-right"},[i("v-uni-view",{staticClass:"text-bold text-black text-center text-xl "},[t._v(t._s(t.numData.death))]),i("v-uni-view",{staticClass:"text-sm text-center margin-top-xs"},[t._v("死亡")]),i("v-uni-view",{staticClass:"text-grey text-sm text-center padding-top-sm"},[t._v("较昨日")]),i("v-uni-view",{staticClass:"margin-top-xs text-black text-center text-xs text-bold"},[t._v(t._s(t.numData.dead_incr>0?"+":"")+t._s(t.numData.dead_incr))])],1),i("v-uni-view",{staticClass:"padding-tb-sm"},[i("v-uni-view",{staticClass:"text-bold text-green text-center text-xl "},[t._v(t._s(t.numData.cure))]),i("v-uni-view",{staticClass:"text-sm text-center margin-top-xs"},[t._v("治愈")]),i("v-uni-view",{staticClass:"text-grey text-sm text-center padding-top-sm"},[t._v("较昨日")]),i("v-uni-view",{staticClass:"margin-top-xs text-green text-center text-xs text-bold"},[t._v(t._s(t.numData.cured_incr>0?"+":"")+t._s(t.numData.cured_incr))])],1)],1),i("v-uni-view",{staticClass:"bg-orange text-center text-content padding-tb-sm"},[t._v("截止时间："+t._s(t.numData.update))])],1),i("v-uni-view",{staticClass:"cu-card margin-0"},[i("v-uni-view",{staticClass:"lee-titles"},[t._v("疫情地图")]),i("v-uni-image",{staticClass:"cu-card radius",staticStyle:{width:"92%"},attrs:{src:t.desc.imgUrl,mode:"widthFix"}})],1)],1):t._e(),i("v-uni-view",{staticClass:"cu-list menu-avatar"},[t._l(e.item,function(e,a){return i("v-uni-view",{key:a,staticClass:"cu-card cu-item",style:"height:"+180*e.length+"upx;",on:{click:function(i){arguments[0]=i=t.$handleEvent(i),t.tz(e.url)}}},[i("v-uni-view",{staticClass:"cu-avatar lg radius",style:"background-image: url("+e.image+");"}),i("v-uni-view",{staticClass:"content"},[i("v-uni-view",{staticClass:"title"},[t._v(t._s(e.title))]),i("v-uni-view",{staticClass:"text-cut"},[t._v(t._s(e.desc))])],1),i("v-uni-view",{staticClass:"action icon-right"})],1)}),e.item.length<1&&0!=a?i("v-uni-view",{staticClass:"cu-card padding"},[i("v-uni-view",{staticClass:"text-center text-content text-gray"},[t._v("暂无内容")])],1):t._e()],2),i("v-uni-view",{staticClass:"cu-card padding"},[i("v-uni-view",{staticClass:"nav-title"},[t._v("免责申明："),i("v-uni-view",{staticClass:"text-content text-xs margin-top-xs"},[t._v("本站所列信息均来自互联网，仅供参考，本站力求但不保证信息的完全准确，如有错漏请以官方媒体披露为准，用户个人对服务的使用承担风险，本站不作任何类型担保！"),i("v-uni-view",{staticClass:"text-xs"},[t._v("By: 后台&数据-祥子"),i("v-uni-text",{staticClass:"text-blue",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.tz("http://werty.cn/")}}},[t._v("werty.cn")]),t._v("前端&设计-李生"),i("v-uni-text",{staticClass:"text-blue",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.tz("https://github.com/admin8756")}}},[t._v("GitHub")])],1),i("v-uni-view",{staticClass:"text-xs",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.tz("https://github.com/wertycn/nCoV")}}},[t._v("GitHub地址:"),i("v-uni-text",{staticClass:"text-blue"},[t._v("https://github.com/wertycn/nCoV")])],1),i("v-uni-view",{staticClass:"text-xs text-bold text-center"},[t._v("© 2019-2020")])],1)],1)],1),i("v-uni-view",{staticClass:"kb flex",staticStyle:{height:"100upx"}})],1)],1)}),1)],1)},n=[];i.d(e,"a",function(){return a}),i.d(e,"b",function(){return n})}}]);