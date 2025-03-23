Here’s a **text-based high-level architecture** of **Consistent Hashing**:  

---

### **🖥️ High-Level Architecture of Consistent Hashing**
```
              +-------------------------------+
              |          Hash Ring            |
              |-------------------------------|
              |      (0)          (100)       |
              |       |              |        |
              |   [VNode A1]     [VNode C1]   |
              |       |              |        |
              | (10)  |              |  (90)  |
              |  |    |              |    |   |
              | [VNode B1]---------[VNode D1] |
              |  (20)  |              |  (80) |
              |  |     |              |    |  |
              | [VNode A2]          [VNode C2]|
              |  (30)  |              |  (70) |
              |  |     |              |    |  |
              | [VNode B2]---------[VNode D2] |
              |  (40)  |              |  (60) |
              |       |              |        |
              |   [VNode A3]     [VNode C3]   |
              |       |              |        |
              | (50)  |              |  (50)  |
              +-------------------------------+

                     (Hash Space from 0 to 100)

```
---

### **📌 Explanation**
1. **Hash Ring Representation**:
   - The circular structure represents the **hash space** (0 to 100).
   - Nodes are placed at specific hash values.

2. **Physical Nodes & Virtual Nodes**:
   - **Physical Nodes (A, B, C, D)** own **multiple virtual nodes** (vNodes) distributed across the ring.
   - Example: `VNode A1, A2, A3` belong to **Physical Node A**.

3. **Request Mapping**:
   - A request (hashed key) is mapped to the nearest **VNode**, which forwards it to the corresponding **physical node**.
   - Example: If a key hashes to **35**, it goes to `VNode A2`, which belongs to **Physical Node A**.

4. **Rebalancing on Node Failure or Addition**:
   - If **Node B fails**, only **its vNodes** (`B1, B2`) are redistributed to nearby nodes (`A, C`).
   - If **a new node (E) joins**, it gets its own vNodes across the ring, **minimizing disruptions**.

---

### **🔥 Key Benefits**
✅ **Even Load Distribution** – Prevents data hotspots.  
✅ **Minimal Rebalancing** – Only affected vNodes need reassignment.  
✅ **Fault Tolerance** – Losing a physical node does not disrupt the entire system.  
✅ **Scalability** – Easily add/remove nodes without massive data movement.  

Would you like further refinements? 🚀