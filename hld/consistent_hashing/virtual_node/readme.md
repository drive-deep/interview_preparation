# 📌 Consistent Hashing with Virtual Nodes (vNodes)

## 🚀 Introduction
Consistent Hashing is a technique used in distributed systems to evenly distribute load across multiple nodes.
However, when nodes are added or removed, **load imbalance** can occur. To solve this, we use **Virtual Nodes (vNodes)**, which spread the load more evenly and improve fault tolerance.

---
## 📖 How It Works
### 1️⃣ **Basic Consistent Hashing (Without vNodes)**
- Each **physical node** is assigned a position on a hash ring.
- Requests (keys) are mapped to the nearest node in the ring.
- **Problem**: If nodes are unevenly placed, some may get overloaded.

**Example:**
```
Hash Ring:
 0 --- A (10) --- B (40) --- C (70) --- 100
```
- If **Node B fails**, all its keys shift to **Node C**, possibly overloading it.
- If a **new node D joins at 50**, many keys are reassigned from **Node C to D**.

---
### 2️⃣ **Consistent Hashing with Virtual Nodes (vNodes)**
Instead of a single physical node in one position, we create **multiple virtual nodes per physical node**, spreading the load more evenly.

**Example:**
```
Hash Ring:
 0 -- A1(10) -- B1(20) -- A2(30) -- C1(40) -- B2(50) -- C2(60) -- A3(70) -- B3(80) -- C3(90) -- 100
```
- **Each physical node (A, B, C) owns multiple virtual nodes**.
- If **Node B fails**, only **its vNodes are reassigned**, not all its data.
- If **Node D joins**, it takes over some vNodes **across the ring**, preventing hotspots.

---
## 📌 Why Use Virtual Nodes?
✅ **Even Load Distribution** – Prevents uneven data distribution.
✅ **Minimal Rebalancing** – Only affected vNodes need reassignment when a node joins/leaves.
✅ **Better Fault Tolerance** – If a node fails, the load is spread across multiple nodes.
✅ **Scalability** – Adding new nodes does not disrupt the system significantly.

---
## 🛠️ Example Implementation (Python)
```python
import hashlib
import bisect

class ConsistentHashing:
    def __init__(self, num_replicas=3):
        self.ring = []
        self.node_map = {}
        self.num_replicas = num_replicas

    def _hash(self, key):
        return int(hashlib.md5(key.encode()).hexdigest(), 16) % 100

    def add_node(self, node):
        for i in range(self.num_replicas):
            vnode_key = f"{node}-{i}"
            h = self._hash(vnode_key)
            bisect.insort(self.ring, h)
            self.node_map[h] = node

    def remove_node(self, node):
        for i in range(self.num_replicas):
            vnode_key = f"{node}-{i}"
            h = self._hash(vnode_key)
            self.ring.remove(h)
            del self.node_map[h]

    def get_node(self, key):
        h = self._hash(key)
        index = bisect.bisect(self.ring, h) % len(self.ring)
        return self.node_map[self.ring[index]]

# Example Usage
ch = ConsistentHashing()
ch.add_node("NodeA")
ch.add_node("NodeB")
ch.add_node("NodeC")

print("Key1 is mapped to:", ch.get_node("Key1"))
print("Key2 is mapped to:", ch.get_node("Key2"))
```

---
## 📌 Real-World Use Cases
🔹 **Distributed Databases** (Amazon DynamoDB, Apache Cassandra)
🔹 **Load Balancers** (NGINX, AWS ELB)
🔹 **Content Delivery Networks** (Cloudflare, Akamai)
🔹 **Distributed Caching** (Redis Cluster, Memcached)

---
## 🏆 Conclusion
Virtual Nodes enhance **scalability, fault tolerance, and load balancing** in consistent hashing. They make distributed systems more **efficient and resilient** to changes.

🚀 **Adopt vNodes for a more balanced, scalable, and fault-tolerant system!**

