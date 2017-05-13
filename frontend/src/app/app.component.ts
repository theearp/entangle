import { Component, OnInit } from '@angular/core';
import { MdSnackBar } from '@angular/material';

import { ShoppingService } from './shared/services/shopping.service';
import { CategoryService } from './shared/services/category.service';
import { MessageService } from './shared/services/message.service';
import { Category, Section } from './shared/models/category';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  categories: Category[];
  sections: Section[];
  cartCount: Number;

  constructor(
    private ss: ShoppingService, 
    private cs: CategoryService,
    private ms: MessageService,
    public snackBar: MdSnackBar
    ) {
      this.ms.get().subscribe(msg => {
        snackBar.open(msg, 'Dismiss');
      })
    }

  ngOnInit() {
     this.cs.getSections()
      .subscribe(
        data => this.sections = data,
        err => this.ms.send(err));

      this.ss.get()
      .subscribe(data => {
        this.cartCount = data.length;
      });
  }
}
