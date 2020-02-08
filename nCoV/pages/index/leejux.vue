<template>
	<!-- 李俊霞的第一个项目 -->
	<view class="page"></view>
</template>

<script>
/* JavaScript */
import std from '../../utils/std.js';
export default {
	data() {
		// 列表渲染条件
		where: {
		}
	},
	methods: {
		/* 页面加载 */
		onLoad(e) {},
		/* 页面触底 */
		onReachBottom(e) {
			/* 页面触底加载的方法 */
			// let page = this.where.page
			// this.[`where.page`] = page+1
			// this.getList()
		},
		/* 页面滚动 */
		onPageScroll(e) {
			// 页面滚动判断
			// if(e.scrollTop<=200){
			// 	this.show = true
			// }else{
			// 	this.show = true
			// }
		},
		/* 页面隐藏 */
		onHide() {},
		/* 页面卸载 */
		onUnload() {},
		/* 页面下拉 */
		onPullDownRefresh() {
			// this.list = [],
			// this.[`where.page`] = 1,
			// this.getList()
		},
		/* 点击了返回按钮 */
		onBackPress() {},
		getList(e) {}
	}
};
</script>

<style>
/* css */
page {
	background: #fff;
}
</style>
