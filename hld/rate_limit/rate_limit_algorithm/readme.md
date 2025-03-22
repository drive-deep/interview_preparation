# **🚀 README: AI-Powered Rate Limiter Microservice (Redis + API Gateway + AI)**
This **Rate Limiter Microservice** efficiently manages API request throttling using **Redis Sorted Set**, integrates **AI for anomaly detection**, and operates behind an **API Gateway**.  

It supports:  
✅ **Fixed Window Counter**  
✅ **Sliding Window Log**  
✅ **Token Bucket Algorithm**  
✅ **Leaky Bucket Algorithm**  
✅ **Redis Sorted Set (Most Optimal)**  
✅ **AI-driven DDoS Detection**  

---

## **📌 High-Level Architecture**
```
      Client Request  
          │  
          ▼  
  ┌───────────────────┐  
  │  API Gateway #1   │  
  └───────────────────┘  
          │  
          ▼  
  ┌───────────────────┐  
  │  Rate Limiter     │  
  │ (Redis + AI)      │  
  └───────────────────┘  
          │  
   Allowed ✅  |  Denied ❌  
          ▼  
  ┌───────────────────┐  
  │  API Gateway #2   │  
  └───────────────────┘  
          │  
          ▼  
  ┌───────────────────┐  
  │ Service Mesh (Istio) │  
  └───────────────────┘  
          │  
          ▼  
  ┌───────────────────┐  
  │ Backend Services  │  
  └───────────────────┘  
```

---

## **📌 Redis Sorted Set (Sliding Window Log)**
This approach:  
- **Uses Redis Sorted Set (ZSET) to store timestamps** of API requests.  
- **Removes old timestamps** outside the window.  
- **Efficiently counts recent requests.**  

### **Implementation**
```python
import time
import redis

# Redis connection
redis_client = redis.Redis(host='localhost', port=6379, decode_responses=True)

def is_rate_limited(user_id: str, max_requests: int, window_seconds: int):
    key = f"rate_limit:{user_id}"
    now = time.time()

    # Start of the sliding window
    window_start = now - window_seconds

    # Remove old requests
    redis_client.zremrangebyscore(key, 0, window_start)

    # Count remaining requests
    request_count = redis_client.zcard(key)

    if request_count >= max_requests:
        return True  # Rate limited

    # Add new request timestamp
    redis_client.zadd(key, {now: now})

    # Set expiry to prevent memory leak
    redis_client.expire(key, window_seconds)
    
    return False  # Allowed

# Example usage
user_id = "123"
if is_rate_limited(user_id, 5, 60):
    print("Too many requests (429)")
else:
    print("Request allowed")
```
### **How It Works**
1️⃣ **Removes old timestamps** beyond the window.  
2️⃣ **Counts active requests** within the window.  
3️⃣ **If limit exceeded, block request (429).**  
4️⃣ **Else, store request & proceed.**  

---

## **📌 All Rate Limiting Algorithms**
| Algorithm | How it Works | Redis Storage | Best For |  
|-----------|-------------|--------------|----------|  
| **Fixed Window** | Counts requests in a fixed interval | `SETEX user_123 60 5` | Basic rate limiting |  
| **Sliding Window Log** | Stores timestamps in a sorted list & trims old ones | `ZADD/ZREMRANGEBYSCORE` | Precise control |  
| **Token Bucket** | Pre-filled tokens consumed per request | `DECR tokens_user_123` | Bursty traffic |  
| **Leaky Bucket** | Requests processed at constant rate | `LPUSH/LPOP queue_user_123` | Smooth traffic |  

---

## **📌 AI-Powered DDoS Detection**
1️⃣ **Extract features:**  
   - Request patterns  
   - Geolocation  
   - Anomaly score  
2️⃣ **Train LSTM or Isolation Forest Model**  
3️⃣ **Flag anomalies & adjust rate limits dynamically**  

### **Database Storage (PostgreSQL)**
```sql
CREATE TABLE ai_ddos_detection (
    id SERIAL PRIMARY KEY,
    user_id TEXT,
    ip_address TEXT,
    request_count INT,
    anomaly_score FLOAT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## **📌 How Data is Stored**
| Redis Key | Value | Expiry |  
|-----------|-------|--------|  
| `rate_limit:user_123` | `{timestamp: timestamp}` | `60s` |  
| `rate_limit:ip_192.0.1` | `{timestamp: timestamp}` | `3600s` |  

---

## **📌 Best Optimal Solution**
✅ **Redis Sorted Set for efficient real-time rate limiting**  
✅ **AI-based DDoS protection for adaptive throttling**  
✅ **Service Mesh for microservices traffic control**  
✅ **API Gateway for security & request routing**  

---

## **📌 API Endpoints**
### **1️⃣ Apply Rate Limit**
```http
POST /rate-limit
{
   "user_id": "123",
   "ip_address": "192.168.1.1"
}
```
- **Checks Redis Sorted Set for active requests.**  
- **If exceeded, return 429 Too Many Requests.**  

---

### **2️⃣ Update Rate Limit Rules**
```http
PUT /rate-limit/rules
{
   "user_id": "123",
   "limit_per_minute": 5
}
```
- **Updates PostgreSQL.**  
- **Syncs new rules to Redis.**  

---

### **3️⃣ AI DDoS Detection**
```http
GET /ddos/anomaly
```
- **Returns flagged users/IPs.**  

---

## **📌 How Requests Flow**
1️⃣ **API Gateway (Ingress) receives request.**  
2️⃣ **Rate Limiter Service (Redis) checks request count.**  
3️⃣ **If allowed, API Gateway forwards to Service Mesh (Istio).**  
4️⃣ **Istio routes request to appropriate microservice.**  
5️⃣ **Response flows back via Service Mesh → API Gateway → Client.**  

---

## **📌 Deployment Considerations**
- **Run Redis in Cluster mode** for scalability.  
- **Use PostgreSQL for persistent rate limit rules.**  
- **Deploy AI model using FastAPI & integrate with Redis.**  
- **Use Envoy/Istio for Service Mesh & rate limit policies.**  

---

### **🚀 Conclusion**
This solution is:  
✅ **Fast** (Redis Sorted Set)  
✅ **Scalable** (Service Mesh & API Gateway)  
✅ **AI-driven** (Detects DDoS attacks)  
✅ **Flexible** (Supports different rate-limiting strategies)  

Would you like a **Docker Compose setup** for deployment? 🚀