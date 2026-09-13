type Node struct {
    key int
    value int
    next *Node
    prev *Node
}

type LRUCache struct {
    capacity int
    cache map[int]*Node
    head *Node
    tail *Node
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        capacity: capacity, 
        cache: make(map[int]*Node),
        head: &Node{},
        tail: &Node{},
    }

    lru.head.next = lru.tail
    lru.tail.prev = lru.head
    return lru
}

func (this *LRUCache) removeNode(node *Node){
    node.next.prev = node.prev
    node.prev.next = node.next
}

func (this *LRUCache) addNode(node *Node){
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) Get(key int) int {
    if node, exist := this.cache[key]; exist{
        this.removeNode(node)
        this.addNode(node)
        return node.value
    }

    return -1
}


func (this *LRUCache) Put(key int, value int) {
    if node, exist := this.cache[key]; exist{
        node.value = value
        this.removeNode(node)
        this.addNode(node)
        return
    }

    newNode := &Node{key: key, value: value}
    this.cache[key] = newNode
    this.addNode(newNode)


    if len(this.cache) > this.capacity{

        lruNode := this.tail.prev
        this.removeNode(lruNode)
        delete(this.cache, lruNode.key)
    }
}


