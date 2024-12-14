

# **Subnetting Explained**

Subnetting is the process of dividing a large network into smaller, more manageable subnetworks. It helps optimize IP address usage, improve network performance, and enhance security.

---

## **What is Subnetting?**
An IP address consists of two parts:
1. **Network Portion**: Identifies the network.
2. **Host Portion**: Identifies devices within the network.

A **subnet mask** defines how the IP address is split into these portions. Subnetting allows you to borrow bits from the host portion to create additional networks.

---

## **Why Subnet?**
- Efficient use of IP addresses.
- Reduces network congestion.
- Enhances security by isolating network segments.
- Simplifies network management.

---

## **How Subnetting Works**

### **Subnet Mask**
The subnet mask determines the number of available hosts and subnets.

| **CIDR Notation** | **Subnet Mask**       | **Hosts per Subnet** | **Number of Subnets (if borrowing bits)** |
|--------------------|-----------------------|-----------------------|-------------------------------------------|
| /24               | 255.255.255.0        | 254                   | -                                         |
| /25               | 255.255.255.128      | 126                   | 2                                         |
| /26               | 255.255.255.192      | 62                    | 4                                         |
| /27               | 255.255.255.224      | 30                    | 8                                         |

---

## **Steps to Calculate Subnets**

### 1. **Determine the Requirements**
   - How many subnets do you need?
   - How many hosts per subnet?

### 2. **Calculate Subnet Mask**
   - Borrow bits from the host portion.
   - Formula:
     - **Number of Subnets**: \( 2^n \) (where \( n \) is the number of bits borrowed).
     - **Hosts per Subnet**: \( 2^h - 2 \) (where \( h \) is the remaining host bits).

### 3. **Determine Network and Broadcast Addresses**
   - **Network Address**: All host bits are `0`.
   - **Broadcast Address**: All host bits are `1`.

---

## **Example: Subnetting a /24 Network**

### **Scenario**:
- Given IP: `192.168.1.0/24`
- Requirement: 4 subnets.

### **Solution**:
1. **Borrow Host Bits**:
   - Original mask: `/24` (255.255.255.0).
   - Borrow 2 bits → New mask: `/26` (255.255.255.192).

2. **Calculate Subnet Ranges**:
   - Subnet increment: \( 256 - 192 = 64 \).

| **Subnet** | **Range**              | **Network Address** | **Broadcast Address** |
|------------|-------------------------|----------------------|------------------------|
| Subnet 1   | 192.168.1.0 - 192.168.1.63  | 192.168.1.0         | 192.168.1.63          |
| Subnet 2   | 192.168.1.64 - 192.168.1.127 | 192.168.1.64        | 192.168.1.127         |
| Subnet 3   | 192.168.1.128 - 192.168.1.191| 192.168.1.128       | 192.168.1.191         |
| Subnet 4   | 192.168.1.192 - 192.168.1.255| 192.168.1.192       | 192.168.1.255         |

3. **Hosts per Subnet**:
   - \( 2^6 - 2 = 62 \) usable hosts.

---

## **Example: Calculating Subnet for /26**

### **Scenario**:
- IP Address: `192.168.10.0/26`.
- Determine the range, network, and broadcast addresses.

### **Solution**:
1. **Subnet Mask**: `/26` (255.255.255.192).
2. **Subnet Increment**: \( 256 - 192 = 64 \).
3. **Calculate Ranges**:

| **Subnet** | **Range**              | **Network Address** | **Broadcast Address** |
|------------|-------------------------|----------------------|------------------------|
| Subnet 1   | 192.168.10.0 - 192.168.10.63  | 192.168.10.0        | 192.168.10.63         |
| Subnet 2   | 192.168.10.64 - 192.168.10.127| 192.168.10.64       | 192.168.10.127        |
| Subnet 3   | 192.168.10.128 - 192.168.10.191| 192.168.10.128      | 192.168.10.191        |
| Subnet 4   | 192.168.10.192 - 192.168.10.255| 192.168.10.192      | 192.168.10.255        |

---

## **Useful Formulas**

1. **Number of Subnets**:
   \[
   2^n
   \]
   Where \( n \) is the number of bits borrowed.

2. **Number of Hosts per Subnet**:
   \[
   2^h - 2
   \]
   Where \( h \) is the remaining host bits.

3. **Subnet Increment**:
   \[
   256 - \text{New Subnet Mask Value}
   \]

---

## **Additional Examples**

### **Subnetting a Class B Network**
- IP: `172.16.0.0/16`
- Requirement: 256 subnets.
- Solution:
  - Borrow 8 bits → `/24` (Subnet Mask: 255.255.255.0).
  - Subnet Range: Each subnet contains 256 addresses, 254 usable hosts.

---

## **Resources**
- Online subnet calculator: [Subnet Calculator](https://www.subnet-calculator.com)
- Learn more about CIDR: [CIDR Notation Guide](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)

---

Let me know if you'd like further clarification or additional examples!
