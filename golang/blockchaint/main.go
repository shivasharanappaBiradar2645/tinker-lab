package main

import (
	"crypto/md5"

	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos          int
	Data         BookCheckout
	Timestamp    string
	Hash         string
	Previoushash string
}

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookCheckout struct {
	BookId    string `json:"book_id"`
	User      string `json:"user"`
	Checkout  string `json:"checkout"`
	IsGenesis bool   `json:"is_genesis"`
}

type BlockChain struct {
	blocks []*Block
}

var blockchain *BlockChain

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := string(b.Pos) + b.Timestamp + string(bytes) + b.Previoushash

	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))

}

func CreateBlock(prevBlock *Block, chechout BookCheckout) *Block {
	block := &Block{}
	block.Previoushash = prevBlock.Hash
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.generateHash()
	return block
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, BookCheckout{IsGenesis: true})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}

func validBlock(block, pb *Block) bool {
	if pb.Hash != block.Previoushash {
		return false
	}
	//skipping pos check and hash check

	return true
}

func (bc *BlockChain) AddBlock(data BookCheckout) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

func newBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not create "))
		return

	}
	h := md5.New()
	io.WriteString(h, book.Title+book.Author)
	book.Id = fmt.Sprintf("%x", h.Sum(nil))
	resp, err := json.MarshalIndent(book, " ", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not create "))
		return

	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func wrtieBook(w http.ResponseWriter, r *http.Request) {
	var checkoutitem BookCheckout

	if err := json.NewDecoder(r.Body).Decode(&checkoutitem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not write"))
		return
	}
	blockchain.AddBlock(checkoutitem)
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockchain.blocks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func main() {
	blockchain = NewBlockChain()
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", wrtieBook).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", r))
}
