<!-- 个人中心页面 -->
<template>
	<view class="page">
		<!-- 切换的菜单 -->
		<scroll-view scroll-x class="bg-white fixed nav margin-top-xs" :class="!show?'fixed':''" scroll-with-animation
		 :scroll-left="scrollLeft">
			<view class="cu-item" :class="index==TabCur?'text-orange cur':''" v-for="(item,index) in tabList" :key="index" @tap="tabSelect"
			 :data-id="index" :data-type='item.id'>
				{{item.title}}
			</view>
		</scroll-view>
		<swiper :indicator-dots="false" :autoplay="false" :interval="1000" :current="TabCur" @change="tabSwiper" style="height: 100vh;width: 100vw;padding-top: 100upx;">
			<swiper-item v-for="(list,index) in allList" :key="index">
				<scroll-view scroll-y="true" style="height: 100vh;width: 100vw;">
					<!-- 条件判断  如果是第一个item -->
					<view v-if="index==0">
						<view class="cu-card padding-0">
							<view class="lee-titles solid-bottom text-content padding-bottom-sm">
								常用导航
							</view>
							<view class="grid col-5 margin-bottom-sm">
								<view class="padding-xs text-center" v-for="(item,index) in menuList" :key="index" @click="tz(item.url)">
									<image class="cu-avatar lg round" :src="item.image" />
									<view class="text-gray text-sm margin-top-xs text-cuts">
										{{item.title}}
									</view>
								</view>
							</view>
						</view>
						<view class="cu-card padding-0">
							<view class="grid col-5">
								<view class="padding-tb-sm solid-right">
									<view class="text-bold text-red text-center text-xl ">
										{{numData.confirmed}}
									</view>
									<view class="text-sm text-center margin-top-xs">
										确诊
									</view>
									<view class="text-grey text-sm text-center padding-top-sm">较昨日</view>
									<view class="margin-top-xs text-red text-center text-xs text-bold">
										{{numData.confirmed_incr>0?'+':''}}{{numData.confirmed_incr}}
									</view>
								</view>
								<view class="padding-tb-sm solid-right">
									<view class="text-bold text-orange text-center text-xl ">
										{{numData.suspected}}
									</view>
									<view class="text-sm text-center margin-top-xs">
										疑似
									</view>
									<view class="text-grey text-sm text-center padding-top-sm">较昨日</view>
									<view class="margin-top-xs text-orange text-center text-xs text-bold">
										{{numData.suspected_incr>0?'+':''}}{{numData.suspected_incr}}
									</view>
								</view>
								<view class="padding-tb-sm solid-right">
									<view class="text-bold text-brown text-center text-xl ">
										{{numData.serious}}
									</view>
									<view class="text-sm text-center margin-top-xs">
										重症
									</view>
									<view class="text-grey text-sm text-center padding-top-sm">较昨日</view>
									<view class="margin-top-xs text-brown text-center text-xs text-bold">
										{{numData.serious_incr>0?'+':''}}{{numData.serious_incr}}
									</view>
								</view>
								<view class="padding-tb-sm solid-right">
									<view class="text-bold text-black text-center text-xl ">
										{{numData.death}}
									</view>
									<view class="text-sm text-center margin-top-xs">
										死亡
									</view>
									<view class="text-grey text-sm text-center padding-top-sm">较昨日</view>
									<view class="margin-top-xs text-black text-center text-xs text-bold">
										{{numData.dead_incr>0?'+':''}}{{numData.dead_incr}}
									</view>
								</view>
								<view class="padding-tb-sm">
									<view class="text-bold text-green text-center text-xl ">
										{{numData.cure}}
									</view>
									<view class="text-sm text-center margin-top-xs">
										治愈
									</view>
									<view class="text-grey text-sm text-center padding-top-sm">较昨日</view>
									<view class="margin-top-xs text-green text-center text-xs text-bold">
										{{numData.cured_incr>0?'+':''}}{{numData.cured_incr}}
									</view>
								</view>
							</view>
							<view class="bg-orange text-center text-content padding-tb-sm">
								截止时间：{{numData.update}}
							</view>
						</view>
						<!-- 地图 -->
						<view class="cu-card margin-0">
							<view class="lee-titles">
								疫情地图
							</view>
							<image class="cu-card radius" :src="desc.imgUrl" mode="widthFix" style="width: 92%;" />
						</view>
						<!-- <image class="cu-card" :src="desc.dailyPic" mode="widthFix" style="width: 92%;" /> -->
					</view>

					<view class="cu-list menu-avatar">
						<view class="cu-card cu-item" @click="tz(item.url)" v-for="(item,idx) in list.item" :key="idx" :style="'height:' + (item.length*180)+'upx;'">
							<view class="cu-avatar lg radius" :style="'background-image: url('+item.image+');'"></view>
							<view class="content">
								<view class="title">
									{{item.title}}
								</view>
								<view class="text-cut">
									{{item.desc}}
								</view>
							</view>
							<view class="action icon-right"></view>
						</view>
						<view class="cu-card padding" v-if="list.item.length<1&&index!=0">
							<view class="text-center text-content text-gray">
								暂无内容
							</view>
						</view>
					</view>
					<view class="cu-card padding">
						<view class="nav-title">
							免责申明：
							<view class="text-content text-xs margin-top-xs">
								本站所列信息均来自互联网，仅供参考，本站力求但不保证信息的完全准确，如有错漏请以官方媒体披露为准，用户个人对服务的使用承担风险，本站不作任何类型担保！
								<view class="text-xs">
									By: 产品&后端 祥子 <text class="text-blue" @click="tz('http://werty.cn/')">werty.cn </text> 前端&设计-李生<text class="text-blue"
									 @click="tz('https://github.com/admin8756')">GitHub </text>
								</view>
								<view class="text-xs" @click="tz('https://github.com/wertycn/nCoV')">
									GitHub地址:
									<text class="text-blue">https://github.com/wertycn/nCoV</text>
								</view>
								<view class="text-xs text-bold text-center">
									© 2019-2020
								</view>
							</view>
						</view>
					</view>
					<view class="kb flex" style="height: 100upx;">
					</view>
				</scroll-view>
			</swiper-item>
		</swiper>
	</view>
	</view>
</template>

<script>
	import std from '../../utils/std.js';
	export default {
		data() {
			return {
				StatusBar: this.StatusBar,
				tabList: [{
					title: '实时数据'
				}],
				menuList: [
					{
						title:'加载中...',
						url:'#',
						image:'http://5b0988e595225.cdn.sohucs.com/images/20170831/ea2692adffd64b82a2d5b8f4d3d8d7d8.gif'
					}
				],
				show: true,
				allList: [{
					item: []
				}],
				desc: {},
				TabCur: 0,
				scrollLeft: 0,
				CustomBar: this.CustomBar,
				numData: {
					confirmed: 0,
					suspected: 0,
					serious:0,
					death: 0,
					cure: 0,
				}
			};
		},
		computed: {
			style() {
				var StatusBar = this.StatusBar;
				var CustomBar = this.CustomBar;
				var bgImage = this.bgImage;
				var style = `height:${CustomBar}px;padding-top:${StatusBar}px;`;
				if (this.bgImage) {
					style = `${style}background-image:url(${bgImage});`;
				}
				return style;
			}
		},
		methods: {
			tabSwiper(e) {
				this.TabCur = e.detail.current;
				this.scrollLeft = (e.detail.current - 1) * 60;
			},
			tabSelect(e) {
				this.TabCur = e.currentTarget.dataset.id;
				this.scrollLeft = (e.currentTarget.dataset.id - 1) * 60;
				// this.where.type = e.currentTarget.dataset.type;
				// this.list = this.allList[e.currentTarget.dataset.type].item
			},
			getMore(e) {
				std.re('original/tx/').then(res => {
					this.desc = res.data.newslist[0].desc
				})
			},
			/* 页面加载 */
			onLoad(e) {
				this.getNum()
				this.getNav()
				this.baidu()
				this.getMore()
				this.getMenu()

			},
			/* 页面触底 */
			onReachBottom(e) {},
			/* 页面滚动 */
			onPageScroll(e) {},
			/* 页面隐藏 */
			onHide() {},
			/* 页面显示 */
			onShow() {},
			/* 页面卸载 */
			onUnload() {},
			/* 页面下拉 */
			onPullDownRefresh() {},
			/* 点击了返回按钮 */
			onBackPress() {},
			/* 获取导航数据 */
			getNav() {
				std.re('nav/app/').then(res => {
					this.allList = std.acc(this.allList, res.data.data)
					let list = res.data.data.map((k, idx) => {
						return {
							title: k.title,
							id: idx
						}
					})
					this.tabList = std.acc(this.tabList, list)
				})
			},
			getMenu() {
				std.re('nav/recommendation').then(res => {
					this.menuList = res.data.data
				})
			},
			/* 获取感染人数 */
			getNum() {
				std.re('latest/tx/').then(res => {
					let data = res.data.data.count
					data.update = std.toTimes(res.data.data.update_time)
					this.numData = data
					this.moreData = res.data.data.info
					let datas = {
						date: std.getYesterdayDate()
					};
					std.re('daily/sum/', datas).then(res => {
						console.log(res)
					})
				})
			},
			/* 百度统计 */
			baidu() {
				var _hmt = _hmt || [];
				(function() {
					var hm = document.createElement("script");
					hm.src = "https://hm.baidu.com/hm.js?cc2f5bf672651df7ddf7d6093ab34350";
					var s = document.getElementsByTagName("script")[0];
					s.parentNode.insertBefore(hm, s);
				})();
			},
			/* 跳转通用方式 */
			goJump(e) {
				console.log(e);
				uni.navigateTo({
					url: e
				});
			},
			tz(e) {
				std.tzWeb(e)
			}
		}
	};
</script>

<style>
	.user-view {
		z-index: -1;
		width: 100%;
		height: 260upx;
	}

	.navigation-list {
		padding-left: 32upx;
		background: #fff;
		margin-top: 36upx;
	}

	.navigation-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding-right: 32upx;
		padding-left: 32upx;
		height: 106upx;
		border-bottom: 1upx solid #f1f1f1;
	}

	.navigation-item .left {
		display: flex;
		align-items: center;
	}

	.navigation-item .right {
		align-items: right;
	}

	.navigation-text {
		margin-left: 21upx;
		font-size: 30upx;
		color: #999;
		font-weight: 500;
		line-height: 38upx;
	}

	.icon-enter {
		width: 14upx;
		height: 24upx;
	}

	.button-null {
		position: relative;
		padding-left: 0upx;
		background: rgba(0, 0, 0, 0);
	}

	.button-hover {
		color: rgba(0, 0, 0, 0.6);
		background-color: #dedede00;
	}

	.button-null::after {
		border: none;
	}

	.mini-img {
		font-size: 42upx;
		color: #888;
	}

	.mini-img2 {
		font-size: 32upx;
		color: #888;
	}

	.text-cuts {
		white-space: nowrap;
		overflow: hidden;
	}
</style>
