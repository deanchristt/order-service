package repository

import (
	"github.com/deanchristt/order-service/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(product entity.Product) entity.Product
	UpdateProduct(product entity.Product) entity.Product
	DeleteProduct(product entity.Product)
	AllProduct() []entity.Product
	FindProductById(productId int) entity.Product
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		db,
	}
}

func (p productConnection) InsertProduct(product entity.Product) entity.Product {
	//TODO implement me
	p.connection.Save(product)
	p.connection.Preload("Customer").Find(&product)
	return product
}

func (p productConnection) UpdateProduct(product entity.Product) entity.Product {
	//TODO implement me
	p.connection.Save(product)
	p.connection.Preload("Customer").Find(&product)
	return product
}

func (p productConnection) DeleteProduct(product entity.Product) {
	//TODO implement me
	p.connection.Delete(product)
}

func (p productConnection) AllProduct() []entity.Product {
	//TODO implement me
	var products []entity.Product
	p.connection.Preload("Customer").Find(&products)
	return products
}

func (p productConnection) FindProductById(productId int) entity.Product {
	//TODO implement me
	var product entity.Product
	p.connection.Preload("Customer").Find(&product, productId)
	return product
}
