package main

var B1 = &Block{Owner: "Carlos", Content: "block 1"}
var Bc = &Blockchain{Root: B1}

type Block struct {
	Owner, Content string
	Next           *Block
}

func NewBlock(owner, content string, next *Block) *Block {
	return &Block{
		Owner:   owner,
		Content: content,
		Next:    next,
	}
}

type Blockchain struct {
	Root *Block
}

func NewBlockchain(root *Block) *Blockchain {
	return &Blockchain{Root: root}
}

func (bc *Blockchain) ListBlocks() []*Block {
	bs := make([]*Block, 0)
	curr := bc.Root
	for {
		if curr != nil {
			bs = append(bs, curr)
			curr = curr.Next
			continue
		}
		break
	}
	return bs
}

func (bc *Blockchain) FindOwner(name string) *Block {
	curr := bc.Root
	for {
		if curr != nil {
			if curr.Owner == name {
				return curr
			}
			curr = curr.Next
			continue
		}
		break
	}
	return nil
}
