# Three-Phase Commit (3PC) Protocol in Distributed Systems

The **Three-Phase Commit (3PC)** protocol is an extension of the **Two-Phase Commit (2PC)** protocol designed to improve fault tolerance and address some of the limitations of 2PC. While 2PC ensures atomicity and consistency in distributed transactions, it can block indefinitely in case of coordinator failure or network partition. 3PC mitigates these issues by introducing a third phase and additional safeguards, ensuring that the system can always make progress even in the presence of failures.

---

## How Three-Phase Commit Works

Three-Phase Commit is divided into three phases:

1. **Phase 1: CanCommit Phase** (Preparation Phase)
2. **Phase 2: PreCommit Phase** (Voting Phase)
3. **Phase 3: Commit Phase** (Decision Phase)

### 1. **CanCommit Phase (Preparation Phase)**
- The **coordinator** sends a `CAN_COMMIT` request to all participants, asking if they are prepared to commit the transaction.
- Each participant replies with either:
  - **`YES`**: The participant is ready to commit.
  - **`NO`**: The participant cannot commit and needs to abort.

### 2. **PreCommit Phase (Voting Phase)**
- If all participants respond with **`YES`** in the **CanCommit** phase, the **coordinator** sends a `PRE_COMMIT` message to all participants. This phase effectively locks the resources and prepares the system for the commit.
- The participants acknowledge receipt of the `PRE_COMMIT` message and prepare to commit the transaction.

### 3. **Commit Phase (Decision Phase)**
- Once all participants acknowledge the `PRE_COMMIT` message, the **coordinator** sends a `COMMIT` message, instructing all participants to finalize the transaction.
- If a failure occurs during the pre-commit phase and some participants fail to respond or acknowledge the `PRE_COMMIT`, the transaction is aborted by the coordinator and all participants are instructed to rollback.

---

## Key Concepts of Three-Phase Commit

- **Coordinator**: The central node that manages the transaction and sends the messages to participants.
- **Participants**: The distributed nodes involved in the transaction, responding to the coordinator’s requests.
- **Atomicity**: Ensures that the transaction is either fully committed or fully rolled back, maintaining consistency across all participants.
- **Non-blocking**: Unlike 2PC, 3PC ensures that participants do not block indefinitely in case of coordinator failure, improving fault tolerance.
- **Fault Tolerance**: 3PC adds an additional layer of protection by introducing a third phase and improving decision-making in the event of network failures or coordinator crashes.

---

## Advantages of Three-Phase Commit

1. **Improved Fault Tolerance**: Unlike 2PC, 3PC is designed to ensure that the system does not block indefinitely. In case of failure, it can continue or abort the transaction safely without getting stuck.
2. **Non-blocking**: The introduction of the **PreCommit** phase allows participants to make decisions more robustly, even in case of failure or network partition, without waiting for the coordinator to recover.
3. **Reduced Risk of Blocking**: The three-phase structure ensures that if the coordinator crashes after the **CanCommit** phase, the participants can recover and decide on the transaction, reducing the risk of blocking.
4. **Progress Guarantees**: 3PC ensures that the transaction can either be completed or safely aborted without getting stuck.

---

## Disadvantages of Three-Phase Commit

1. **Complexity**: Compared to 2PC, 3PC introduces an additional phase and more message exchanges, which increases the complexity of implementation and transaction management.
2. **Performance Overhead**: The added phase and additional communication between the coordinator and participants can result in higher latency and more network overhead.
3. **Limited by Partial Failures**: While 3PC improves fault tolerance, it is still not immune to certain failure scenarios, especially when there are network partitions between the coordinator and the participants.
4. **Not Perfect for Every Scenario**: Some systems, especially those with high latency or a large number of participants, might not benefit as much from the added phase.

---

## Example Workflow of Three-Phase Commit

### Scenario 1: Successful Commit

1. **Coordinator** sends `CAN_COMMIT` to all participants.
2. **Participant 1**: Responds with **`YES`**.
3. **Participant 2**: Responds with **`YES`**.
4. **Coordinator** sends `PRE_COMMIT` to all participants.
5. **Participant 1**: Acknowledges **`PRE_COMMIT`**.
6. **Participant 2**: Acknowledges **`PRE_COMMIT`**.
7. **Coordinator** sends `COMMIT` to all participants.
8. **Participants** commit the transaction.

### Scenario 2: Failure During Pre-Commit Phase

1. **Coordinator** sends `CAN_COMMIT` to all participants.
2. **Participant 1**: Responds with **`YES`**.
3. **Participant 2**: Responds with **`YES`**.
4. **Coordinator** sends `PRE_COMMIT` to all participants.
5. **Participant 1**: Acknowledges **`PRE_COMMIT`**.
6. **Participant 2** fails to acknowledge due to a crash.
7. **Coordinator** aborts the transaction, instructs participants to rollback.

---

## Real-World Use Cases of Three-Phase Commit

1. **Distributed Databases**: 3PC is used to ensure consistency across multiple database nodes in a distributed database setup, where the transactions span multiple systems.
2. **Microservices Transactions**: When multiple microservices need to coordinate on a distributed transaction, 3PC can help ensure that the system avoids deadlock and progresses safely.
3. **Financial Transactions**: In high-stakes scenarios like financial transactions or e-commerce payments, 3PC ensures that transactions are atomic, consistent, and can handle failures gracefully.
4. **Distributed File Systems**: In systems where files are being updated or transferred across multiple nodes, 3PC ensures data consistency even in the face of failures.

---

## Three-Phase Commit: A Beautiful README

---

# Three-Phase Commit (3PC) Protocol in Distributed Systems

The **Three-Phase Commit (3PC)** protocol is an enhancement over the classic **Two-Phase Commit (2PC)**, designed to improve fault tolerance and prevent blocking in distributed transactions. 3PC helps ensure that distributed systems can reach a consensus and maintain **atomicity**, **consistency**, and **durability** (ACID properties) even in the presence of network partitions or system failures.

---

## How It Works

Three-Phase Commit introduces a third phase, **PreCommit**, between the **CanCommit** and **Commit** phases to prevent blocking and ensure progress.

### Phases:

1. **CanCommit Phase**: Coordinator asks participants if they are ready to commit.
2. **PreCommit Phase**: If all participants agree, the coordinator sends a `PRE_COMMIT` message to lock resources.
3. **Commit Phase**: The coordinator sends a `COMMIT` message, and all participants finalize the transaction.

---

## Advantages

- **Improved Fault Tolerance**: Ensures that the system doesn't block indefinitely during failures.
- **Non-blocking**: Progresses even if the coordinator fails or participants get disconnected.
- **Reduced Risk of Blocking**: Safeguards against indefinite blocking during transaction management.

---

## Disadvantages

- **Increased Complexity**: More phases and messages, leading to more complex implementation.
- **Performance Overhead**: Additional phase can introduce network delays and higher latency.
- **Limited by Network Partitions**: Still vulnerable to certain types of network failures.

---

## Real-World Use Cases

- **Distributed Databases**
- **Microservices Transaction Management**
- **Financial Transactions**
- **Distributed File Systems**

---

## Conclusion

The **Three-Phase Commit (3PC)** protocol is a powerful mechanism for ensuring fault tolerance and consistency in distributed transactions. It solves many of the problems associated with **Two-Phase Commit (2PC)**, such as blocking and failure recovery, making it a great choice for systems where progress and consistency are critical, even in the face of network or system failures.

---

This README provides an in-depth overview of the **Three-Phase Commit (3PC)** protocol, including its phases, advantages, limitations, and practical applications in distributed systems.
