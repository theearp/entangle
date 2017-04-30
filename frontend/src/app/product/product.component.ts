import { Component, OnInit } from '@angular/core';

import { Product } from '../shared/models/product';
import { ProductService } from '../shared/services/product.service';
import { ShoppingService } from '../shared/services/shopping.service';

@Component({
  selector: 'product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.css'],
  providers: [ProductService]
})
export class ProductComponent implements OnInit {
  products: Product[];
  selectedProduct: Product;
  constructor(
    private ps: ProductService, 
    private ss: ShoppingService
    ) {
  }
  
  ngOnInit() {
    this.ps.getProducts()
    .subscribe(data => this.products = data)
  }

  updateShoppingCart(product: Product) {
    this.ss.addItemToCart(product);
    console.log(product);
  }

  showDetails(product: Product) {
    this.selectedProduct = product;
  }
}
