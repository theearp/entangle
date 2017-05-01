import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params }   from '@angular/router';
import { Location }                 from '@angular/common';

import { Product } from '../shared/models/product';
import { ProductService } from '../shared/services/product.service';
import { ShoppingService } from '../shared/services/shopping.service';

import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'product-category',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.css'],
  providers: [ProductService]
})

export class ProductCategoryComponent implements OnInit {
  products: Product[];
  selectedProduct: Product;
  constructor(
    private ps: ProductService, 
    private ss: ShoppingService,
    private route: ActivatedRoute
    ) {
  }
  
  ngOnInit() {
    this.route.params
      .switchMap((params: Params) => this.ps.getProductsByCategory(params['id']))
      .subscribe(data => this.products = data);
  }

  updateShoppingCart(product: Product) {
    this.ss.addItemToCart(product);
    console.log(product);
  }

  showDetails(product: Product) {
    this.selectedProduct = product;
  }
}