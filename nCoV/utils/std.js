/**
 * 
 * LeeTool - JS - Uni版本
 * 低效的 - 自定义常用方法封装库
 * 更新时间： 2019 年11月15日16: 23: 23
 * 这是一个懒人使用的js工具， 大部分的方法都是配置好的。 有一些都不需要传任何参数。
 * 我们不追求很高的自定义。 我们只为了简单高效的写外包项目。 所以这是一个wbjs。
 * 请先引用这个js。 才能使用
 * 引入方法:
 * import std from "../../utils/std.js" 
 * 
 **/

/* 静态类 */
const host = "https://ncov-api.werty.cn:2021/"

/* 请求封装 */
function request(url, datas, type) {
	return new Promise((resolve, reject) => {
		uni.request({
			url: host + url,
			data: datas || {},
			method: type || 'GET',
			// header: {
			// 	"Content-Type": "application/x-www-form-urlencoded"
			// },
			success: res => {
				// console.log(url, res)
				resolve(res);
				// if (res.data.code == 0) resolve(res);
				// else toast(res.data.msg)
			},
			fail: res => {
				console.error(res.data)
				reject(res.data);
			},
			error: e => {
				return reject('网络出错');
			}
		})
	});
}

/* 假的软件更新 */
function updataBtn() {
	let sjs = randomNum(1, 12)
	uni.showLoading({
		title: '更新中...'
	})
	setTimeout(function() {
		uni.hideLoading()
		uni.showModal({
			title: '温馨提醒',
			content: '您现在已是最新版本！',
			showCancel: false,
			confirmText: '我知道了',
		});
	}, sjs * 579);
}
/* 获取昨天的时间 */
function getYesterdayDate() {
    let today = new Date();
    today.setTime(today.getTime() - 24 * 60 * 60 * 1000);
    let month = today.getMonth() + 1;
    if (month < 10) {
        month = "0" + month
    }
    let day = today.getDate();
    if (day < 10) {
        day = "0" + day
    }
    return today.getFullYear() + "-" + month + "-" + day;
}
/*  假的清除缓存 */
function clearCache() {
	toast('清除成功')
}

/* 显示一个简单的弹窗 */

function toast(e) {
	uni.showToast({
		title: e || '错误',
		duration: 2000,
		icon: 'none'
	})
}

/* 图片上传接口封装 ,自动返回图片上传好的地址.*/
function upDataImage() {
	return new Promise((resolve, reject) => {
		if (uni.getStorageSync('token')) {
			uni.chooseImage({
				success: (chooseImageRes) => {
					const tempFilePaths = chooseImageRes.tempFilePaths;
					const uploadTask = uni.uploadFile({
						url: host + 'common/upload',
						filePath: tempFilePaths[0],
						name: 'file',
						formData: {},
						success: (ress) => {
							toast('上传成功')
							resolve({
								img: JSON.parse(ress.data).data.url,
								url: chooseImageRes.tempFilePaths[0]
							})
						},
						fail: (err) => {
							toast('上传失败')
							console.error(err)
							reject: err
						},
					});
					uploadTask.onProgressUpdate((ret) => {
						console.log(ret)
						// 测试条件，取消上传任务。
						// if (ret.progress > 50) {
						// 	uploadTask.abort();
						// }
					});
				}
			});
		} else {
			toast('您还没有登录')
			reject({
				data: '未登录!'
			})
		}
	})
}

/* 跳转方法封装 */
function tz(e) {
	if (e) {
		uni.navigateTo({
			url: e
		})
	} else {
		toast('参数错误')
	}
}

/* 切换tabbar ,
   传值tabbar的路径,
   可以跳转.否则跳首页.
 */
function sw(e) {
	if (e) {
		uni.switchTab({
			url: e,
		})
	} else {
		uni.switchTab({
			url: "/pages/index/index.vue",
		})
	}
}

/* 拼接数据，因为老是记不住这个 */

function acc(arr1, arr2) {
	if (arr2) {
		return (arr1 || []).concat(arr2 || [])
	} else {
		toast('没有更多了...')
	}
}

/* 返回页面
   如果传递数字则返回指定数字层数
   如果没有传数字。默认返回上一层
 */

function back(e) {
	if (e)
		uni.navigateBack({
			delta: e
		})
	else
		uni.navigateBack({
			delta: 1
		})
}

/* 动态设置标题栏 */
function setNavTitle(e) {
	uni.setNavigationBarTitle({
		title: e || ''
	})
}

/* 延迟3秒后退出 */
function sellpBcak() {
	setTimeout(function() {
		back()
	}, 3 * 579);
}

/* 产生随机数 
   传数字。传最小，传最大，
   别问为什么，反正百度的 
*/
function randomNum(minNum, maxNum) {
	switch (arguments.length) {
		case 1:
			return parseInt(Math.random() * minNum + 1, 10);
			break;
		case 2:
			return parseInt(Math.random() * (maxNum - minNum + 1) + minNum, 10);
			break;
		default:
			return 0;
			break;
	}
}

/* 显示一个模态弹窗 */
function tips(msg,title,cancel,cbtn){
	uni.showModal({
		title:title||'温馨提醒',
		content:msg||'错误',
		showCancel:cancel||false,
		confirmText:cbtn||'我知道了'
	})
}

/* 获取时间戳*/
function getTime() {
	return Date.parse(new Date());
}

/* 关闭所有页面到一个新的页面 */
function rel(e) {
	uni.reLaunch({
		url: e
	})
}

/* 输入时间戳,转日期*/
function toTime(e) {
	if (!e) {
		var date = new Date()
	} else if(e < 1546272000000) {
		var date = new Date(e * 1000);
	}else{
		var date = new Date(e);
	}
	let [y, m, d] = [date.getFullYear(), date.getMonth() + 1, date.getDate()]
	m = m < 10 ? ("0" + m) : m;
	d = d < 10 ? ("0" + d) : d;
	return y + "-" + m + "-" + d;
}

/* 输入时间戳,转日期*/
function toMiniTime(e) {
	if (e < 1546272000000) {
		var date = new Date(e * 1000);
	} else {
		var date = new Date(e);
	}
	let [y, m, d] = [date.getFullYear(), date.getMonth() + 1, date.getDate()]
	let [h ,mm ,s] =[date.getHours(),date.getMinutes(),date.getSeconds()]
	m = m < 10 ? ("0" + m) : m;
	d = d < 10 ? ("0" + d) : d;
	h = h < 10 ? ("0" + h) : h;
	mm = mm < 10 ? ("0" + mm) : mm;
	s = s < 10 ? ("0" + s) : s;
	return m + "月" + d +'日'+ h +':'+mm;
}

/* 跳转网页,直接传链接. h5 app 小程序通用*/
function tzWeb(e){
	/* app的方法 */
	//#ifdef APP-PLUS
	plus.runtime.openURL(e, function(res) {
		console.log(res);
	});
	//#endif
	
	/*  h5的方法 */
	//#ifdef H5
	window.location.href = e
	//#endif
	/* 小程序的方法 */
	//#ifndef H5
	tz('news?id=' + e)
	//#endif
}

/* 输入时间戳,转日期时间*/
function toTimes(e) {
	if (e < 1546272000000) {
		var date = new Date(e * 1000);
	} else {
		var date = new Date(e);
	}
	let [y, m, d] = [date.getFullYear(), date.getMonth() + 1, date.getDate()]
	let [h ,mm ,s] =[date.getHours(),date.getMinutes(),date.getSeconds()]
	m = m < 10 ? ("0" + m) : m;
	d = d < 10 ? ("0" + d) : d;
	h = h < 10 ? ("0" + h) : h;
	mm = mm < 10 ? ("0" + mm) : mm;
	s = s < 10 ? ("0" + s) : s;
	return y + "-" + m + "-" + d +' '+ h +':'+mm+':'+s;
}

export default {
	updataBtn,
	toTimes,
	upDataImage,
	getTime,
	tzWeb,
	getYesterdayDate,
	toTime,
	acc,
	tz,
	tips,
	rel,
	sw,
	setNavTitle,
	clearCache,
	toast,
	re: request,
	back,
	toMiniTime,
	sellpBcak,
	randomNum
}
