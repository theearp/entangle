import { Component } from '@angular/core';

import { ShoppingService } from './shared/services/shopping.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  constructor(private ss: ShoppingService) {}

  showCart() {
    console.log('Number of items in cart: ' + this.ss.cartItems.length);
  }
}
