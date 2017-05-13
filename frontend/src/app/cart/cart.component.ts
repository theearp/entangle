import { Component, OnInit } from '@angular/core';

import { ShoppingService } from '../shared/services/shopping.service';
import { Listing } from '../shared/models/listing';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.css']
})
export class CartComponent implements OnInit {
  cart: Listing[];

  constructor(private shoppingService: ShoppingService) { }

  ngOnInit() {
    this.shoppingService.get()
    .subscribe(data => this.cart = data);
  }
}
