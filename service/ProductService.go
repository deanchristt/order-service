package service

//
//import (
//	"fmt"
//	"github.com/deanchristt/order-service/dto"
//	"github.com/deanchristt/order-service/entity"
//	"github.com/deanchristt/order-service/repository"
//	"github.com/mashingan/smapping"
//	"log"
//)
//
//type ProductService interface {
//	Insert(dto dto.ProductCreateDto) entity.Product
//	Update(dto dto.ProductUpdateDto) entity.Product
//	Delete(product entity.Product)
//	All() []entity.Product
//	FindById(productId int) entity.Product
//	IsAllowedToEdit(customerId string, productId int) bool
//}
//
//type productService struct {
//	productRepository repository.ProductRepository
//}
//
//func (p productService) Insert(dto dto.ProductCreateDto) entity.Product {
//	product := entity.Product{}
//	err := smapping.FillStruct(&product, smapping.MapFields(dto))
//	if err != nil {
//		log.Fatalf("Field map %v", err)
//	}
//	res := p.productRepository.InsertProduct(product)
//	return res
//}
//
//func (p productService) Update(dto dto.ProductUpdateDto) entity.Product {
//	product := entity.Product{}
//	err := smapping.FillStruct(&product, smapping.MapFields(dto))
//	if err != nil {
//		log.Fatalf("Field map %v", err)
//	}
//	res := p.productRepository.UpdateProduct(product)
//	return res
//}
//
//func (p productService) Delete(product entity.Product) {
//	p.productRepository.DeleteProduct(product)
//}
//
//func (p productService) All() []entity.Product {
//	return p.productRepository.AllProduct()
//}
//
//func (p productService) FindById(productId int) entity.Product {
//	return p.productRepository.FindProductById(productId)
//}
//
//func (p productService) IsAllowedToEdit(customerId string, productId int) bool {
//	product := p.productRepository.FindProductById(productId)
//	id := fmt.Sprintf("%v", product.CustomerId)
//	return customerId == id
//}
//
//func NewProductService(productRepo repository.ProductRepository) ProductService {
//	return &productService{
//		productRepository: productRepo,
//	}
//}
