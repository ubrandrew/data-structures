class LinkedList:
    def __init__(self) :
        self.head = None
    
    def __str__(self):
        string = ""
        curr = self.head
        while curr != None:
            string += str(curr.value) + " -> "
            curr = curr.next

        return string

    def append(self, val):
        newNode = Node(val)
        if self.head == None:
            self.head = newNode
            return
        curr = self.head
        while curr.next != None:
            curr = curr.next
        
        curr.next = newNode
        return
    


class Node: 
    def __init__(self, val):
        self.next = None
        self.value = val

if __name__ == "__main__":
    linkedlist = LinkedList()
    linkedlist.append(1)
    linkedlist.append("ASD")
    print(linkedlist)