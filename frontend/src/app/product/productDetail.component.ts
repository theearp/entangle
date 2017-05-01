import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute, Params }   from '@angular/router';
import { Location }                 from '@angular/common';
import { ProductService } from '../shared/services/product.service';
import { ShoppingService } from '../shared/services/shopping.service';
import { Product } from '../shared/models/product';

import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'product-detail',
  templateUrl: './productDetail.component.html',
  styleUrls: ['./productDetail.component.css'],
  providers: [ProductService]
})
export class ProductDetailComponent implements OnInit {
  product: Product;

  constructor(
    private ps: ProductService,
    private route: ActivatedRoute,
    private location: Location,
    private ss: ShoppingService
  ) {}

  ngOnInit() {
    this.route.params
      .switchMap((params: Params) => this.ps.getProduct(params['id']))
      .subscribe(data => this.product = data);
  }

  updateShoppingCart(product) {
    this.ss.addItemToCart(product);
  }
}
