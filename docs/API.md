# 奶茶点餐系统 API 文档

## 基本信息

- **项目名称**: 奶茶点餐系统 (Milktea Ordering App)
- **API 基础 URL**: `http://localhost:8080/api`
- **内容类型**: `application/json`

---

## 1. 奶茶管理接口

### 1.1 获取所有奶茶

获取所有奶茶产品列表，支持分页和分类筛选。

**请求方法**: `GET`

**接口路径**: `/api/teas`

**查询参数**:

| 参数名 | 类型 | 必需 | 说明 | 示例 |
|--------|------|------|------|------|
| `category` | string | 否 | 奶茶分类 | `经典奶茶` |
| `page` | integer | 否 | 页码（最小值为1） | `1` |
| `page_size` | integer | 否 | 每页数量（1-100，默认10） | `10` |

**示例请求**:
```bash
curl -X GET "http://localhost:8080/api/teas?category=经典奶茶&page=1&page_size=10"
```

**成功响应** (200 OK):
```json
{
  "total": 20,
  "page": 1,
  "page_size": 10,
  "items": [
    {
      "id": 1,
      "name": "珍珠奶茶",
      "price": 12.5,
      "category": "经典奶茶",
      "description": "经典珍珠奶茶，口感醇厚",
      "stock": 100,
      "image_url": "http://example.com/tea1.jpg"
    },
    {
      "id": 2,
      "name": "草莓奶茶",
      "price": 14.0,
      "category": "水果奶茶",
      "description": "新鲜草莓搭配奶茶",
      "stock": 50,
      "image_url": "http://example.com/tea2.jpg"
    }
  ]
}
```

---

### 1.2 获取奶茶详情

根据 ID 获取单个奶茶产品的详细信息。

**请求方法**: `GET`

**接口路径**: `/api/teas/:id`

**路径参数**:

| 参数名 | 类型 | 必需 | 说明 |
|--------|------|------|------|
| `id` | integer | 是 | 奶茶产品 ID |

**示例请求**:
```bash
curl -X GET "http://localhost:8080/api/teas/1"
```

**成功响应** (200 OK):
```json
{
  "id": 1,
  "name": "珍珠奶茶",
  "price": 12.5,
  "category": "经典奶茶",
  "description": "经典珍珠奶茶，口感醇厚",
  "stock": 100,
  "image_url": "http://example.com/tea1.jpg"
}
```

**错误响应** (404 Not Found):
```json
{
  "error": "奶茶不存在"
}
```

---

### 1.3 创建奶茶

创建新的奶茶产品（需要管理员权限）。

**请求方法**: `POST`

**接口路径**: `/api/teas`

**请求头**:
```
Content-Type: application/json
```

**请求体**:

| 字段名 | 类型 | 必需 | 说明 |
|--------|------|------|------|
| `name` | string | 是 | 奶茶名称 |
| `price` | float | 是 | 价格（单位：元） |
| `category` | string | 是 | 分类 |
| `description` | string | 否 | 描述信息 |
| `stock` | integer | 是 | 库存数量 |
| `image_url` | string | 否 | 图片 URL |

**示例请求**:
```bash
curl -X POST "http://localhost:8080/api/teas" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "芒果奶茶",
    "price": 13.5,
    "category": "水果奶茶",
    "description": "新鲜芒果与奶茶的完美结合",
    "stock": 80,
    "image_url": "http://example.com/tea3.jpg"
  }'
```

**成功响应** (201 Created):
```json
{
  "id": 3,
  "name": "芒果奶茶",
  "price": 13.5,
  "category": "水果奶茶",
  "description": "新鲜芒果与奶茶的完美结合",
  "stock": 80,
  "image_url": "http://example.com/tea3.jpg"
}
```

---

### 1.4 更新奶茶

更新现有的奶茶产品信息（需要管理员权限）。

**请求方法**: `PUT`

**接口路径**: `/api/teas/:id`

**路径参数**:

| 参数名 | 类型 | 必需 | 说明 |
|--------|------|------|------|
| `id` | integer | 是 | 奶茶产品 ID |

**请求体**: 同创建奶茶接口，所有字段都可选

**示例请求**:
```bash
curl -X PUT "http://localhost:8080/api/teas/1" \
  -H "Content-Type: application/json" \
  -d '{
    "price": 13.0,
    "stock": 120
  }'
```

**成功响应** (200 OK):
```json
{
  "id": 1,
  "name": "珍珠奶茶",
  "price": 13.0,
  "category": "经典奶茶",
  "description": "经典珍珠奶茶，口感醇厚",
  "stock": 120,
  "image_url": "http://example.com/tea1.jpg"
}
```

---

### 1.5 删除奶茶

删除指定的奶茶产品（需要管理员权限）。

**请求方法**: `DELETE`

**接口路径**: `/api/teas/:id`

**路径参数**:

| 参数名 | 类型 | 必需 | 说明 |
|--------|------|------|------|
| `id` | integer | 是 | 奶茶产品 ID |

**示例请求**:
```bash
curl -X DELETE "http://localhost:8080/api/teas/3"
```

**成功响应** (204 No Content):
```
(空响应体)
```

**错误响应** (404 Not Found):
```json
{
  "error": "奶茶不存在"
}
```

---

## 2. 订单管理接口

### 2.1 创建订单

创建新的订单，支持同时购买多个产品。

**请求方法**: `POST`

**接口路径**: `/api/orders`

**请求头**:
```
Content-Type: application/json
```

**请求体**:

| 字段名 | 类型 | 必需 | 说明 |
|--------|------|------|------|
| `user_id` | integer | 是 | 用户 ID |
| `items` | array | 是 | 订单项数组（最少1项） |
| `items[].product_id` | integer | 是 | 产品 ID |
| `items[].quantity` | integer | 是 | 购买数量（最少1） |

**示例请求**:
```bash
curl -X POST "http://localhost:8080/api/orders" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "items": [
      {
        "product_id": 1,
        "quantity": 2
      },
      {
        "product_id": 2,
        "quantity": 1
      }
    ]
  }'
```

**成功响应** (201 Created):
```json
{
  "order_id": "ORD20231209001",
  "status": "pending",
  "total_price": 39.0,
  "message": "订单创建成功，已发送到Kafka消息队列",
  "created_at": "2023-12-09T10:30:45Z"
}
```

**订单状态说明**:

| 状态值 | 状态 | 说明 |
|--------|------|------|
| `0` | pending | 已下单 / 待制作 |
| `1` | preparing | 制作中 |
| `2` | completed | 制作完成 / 待取餐 |
| `3` | cancelled | 已取消 |

**错误响应** (400 Bad Request):
```json
{
  "error": "缺少必要字段或参数验证失败"
}
```

**错误响应** (404 Not Found):
```json
{
  "error": "用户或产品不存在"
}
```

---

## 3. 错误响应格式

所有 API 错误响应都遵循以下格式：

```json
{
  "error": "错误描述信息",
  "code": "错误代码（可选）",
  "timestamp": "2023-12-09T10:30:45Z"
}
```

### 常见HTTP状态码

| 状态码 | 说明 |
|--------|------|
| `200` | 请求成功 |
| `201` | 资源创建成功 |
| `204` | 请求成功（无返回内容） |
| `400` | 请求参数错误 |
| `404` | 资源不存在 |
| `500` | 服务器内部错误 |

---

## 4. 技术特性

### 4.1 消息队列集成

- **消息队列**: Apache Kafka
- **用途**: 订单创建后，系统会自动将订单信息发送到 Kafka 消息队列
- **消费者**: 后端消费服务监听订单消息，实时更新订单状态（制作中 → 完成）

### 4.2 数据库设计

- **ORM框架**: GORM
- **关联关系**:
  - Order 与 User 的多对一关系
  - Order 与 Product 的多对一关系

---

## 5. 使用示例

### 完整的订单流程

1. **获取奶茶列表**
   ```bash
   curl -X GET "http://localhost:8080/api/teas?page=1&page_size=5"
   ```

2. **获取奶茶详情**
   ```bash
   curl -X GET "http://localhost:8080/api/teas/1"
   ```

3. **创建订单**
   ```bash
   curl -X POST "http://localhost:8080/api/orders" \
     -H "Content-Type: application/json" \
     -d '{
       "user_id": 1,
       "items": [
         {"product_id": 1, "quantity": 2}
       ]
     }'
   ```

---

## 6. 注意事项

- 所有 API 响应时间通常在 100-500ms 内
- 请确保在创建订单时提供的 `user_id` 和 `product_id` 都存在于系统中
- 产品库存实时更新，请注意库存充足后再下单
- Kafka 消费者异步更新订单状态，可能有数秒延迟
- 建议使用 HTTP 连接池以提高性能

---

**最后更新**: 2025年12月9日
