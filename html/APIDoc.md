# API 文档
## 接口内容 
1. 病患同行车次寻人

返回数据结构
```js
    {
      "code": 0,
      "msg": "",
      "data": [
        {
          "id": 46,
          // 日期
          "date": "2020-01-26", 
          // 开始时间
          "start_time": "2020-01-25T10:45:00.000Z",
          // 结束时间
          "end_time": "2020-01-25T16:20:00.000Z",
          // 交通类型
          "type": 1,
          // 车次代码
          "train_number": "TR134",
          // 备注
          "remark": "",
          // 车厢
          "crriage_number": "",
          // 始发站
          "start_station": "新加坡",
          // 终点站
          "end_station": "西安",
          // 数据来源
          "source": "https://mp.weixin.qq.com/s/jEBlMLCgSTD9AcpcV1yRtg",
          // 发布媒体
          "publish_media": "央视新闻",
          // 是否验证
          "verified": 1,
          // 创建时间
          "create_time": "2020-01-27T09:54:12.000Z",
          // 更新时间
          "update_time": "2020-01-27T09:54:12.000Z"
        }
      ]
    }
```