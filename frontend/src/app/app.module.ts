import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';
import { AppRoutingModule } from './app.routing';
import { MaterialModule, MdIconRegistry } from '@angular/material';
import { NglModule } from 'ng-lightning/ng-lightning';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


import { AppComponent } from './app.component';
import { ProductComponent } from './product.component';
import { ProductDetailComponent } from './productDetail.component';
import { ShoppingService } from './shopping.service';
import { ListingComponent } from './listing/listing.component';

@NgModule({
  declarations: [
    AppComponent,
    ProductComponent,
    ProductDetailComponent,
    ListingComponent
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
    ShoppingService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
