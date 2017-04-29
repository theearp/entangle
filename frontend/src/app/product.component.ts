import { Component, OnInit } from '@angular/core';

import { Listing } from './product';
import { ProductService } from './product.service';
import { ShoppingService } from './shopping.service';

@Component({
  selector: 'product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.css'],
  providers: [ProductService]
})
export class ProductComponent implements OnInit {
  listings: Listing[];
  constructor(
    private ps: ProductService, 
    private ss: ShoppingService
    ) {
  }
  
  ngOnInit() {
    this.ps.getPopularListings()
    .subscribe(data => this.listings = data)
  }

  updateShoppingCart(listing: Listing) {
    this.ss.addItemToCart(listing);
    console.log(listing);
  }
}
