class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.next = None

class HashMap:
    def __init__(self, capacity: int = 1000):
        self.capacity = capacity        
        self.buckets = [None] * capacity

    def _hash(self, key: int):
        return key % self.capacity

    def put(self, key: int, value):
        index = self._hash(key)
        current = self.buckets[index]
        if not current:
            self.buckets[index] = Node(key, value)
            return
        while current:
            if current.key == key:
                current.value = value
                return
            elif not current.next:
                break
            else:
                current = current.next
                continue
        current.next = Node(key, value)    
          

    def get(self, key: int):
        index = self._hash(key)
        if self.buckets[index]:
            current = self.buckets[index]
            while current:
                if current.key == key:
                    return current.value
                else:
                    current = current.next
        return None

    def remove(self, key: int):
        index = self._hash(key)
        if not self.buckets[index]:
            print(f"key {key} does not exist")
        else:
            current = self.buckets[index]
            if current.key == key:
                self.buckets[index] = current.next      

            prev = current
            current = current.next
            while current:
                if current.key == key:
                    prev.next = current.next
                    return
                else:
                    prev = current
                    current = current.next


hash_map = HashMap(20)
for i in range(100):
    hash_map.put(i, f"value_{i}")


print(hash_map.get(6))
print(hash_map.get(86))
print(hash_map.get(101))
hash_map.put(86 , f"updated_value_{86}")
print(hash_map.get(86))
hash_map.remove(86)
print(hash_map.get(86))