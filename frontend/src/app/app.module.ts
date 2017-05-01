import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';
import { AppRoutingModule } from './app.routing';
import { MaterialModule, MdIconRegistry } from '@angular/material';
import { NglModule } from 'ng-lightning/ng-lightning';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


import { AppComponent } from './app.component';
import { ProductComponent } from './product/product.component';
import { ProductDetailComponent } from './product/productDetail.component';
import { PopularProductsComponent } from './product/popular.component';
import { ProductCategoryComponent } from './product/productCategory.component';
import { SharedComponent } from './shared/shared.component';
import { CartComponent } from './cart/cart.component';

import { ShoppingService } from './shared/services/shopping.service';
import { CategoryService } from './shared/services/category.service';

@NgModule({
  declarations: [
    AppComponent,
    ProductComponent,
    ProductDetailComponent,
    PopularProductsComponent,
    ProductCategoryComponent,
    SharedComponent,
    CartComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    JsonpModule,
    MaterialModule,
    AppRoutingModule,
    NglModule.forRoot(),
    BrowserAnimationsModule
  ],
  providers: [
    MdIconRegistry, 
    ShoppingService,
    CategoryService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
