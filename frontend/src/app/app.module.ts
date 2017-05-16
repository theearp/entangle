import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';
import { AppRoutingModule } from './app.routing';
import { MaterialModule, MdIconRegistry } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


import { AppComponent } from './app.component';
import { SharedComponent } from './shared/shared.component';
import { CartComponent } from './cart/cart.component';

import { ShoppingService } from './shared/services/shopping.service';
import { CategoryService } from './shared/services/category.service';
import { ListingService } from './shared/services/listing.service';
import { MessageService } from './shared/services/message.service';
import { ImageService } from './shared/services/image.service';
import { AboutComponent } from './about/about.component';
import { ListingComponent } from './listing/listing.component';
import { ListingDetailComponent } from './listing/listing-detail.component';
import { ImageComponent } from './image/image.component';

@NgModule({
  declarations: [
    AppComponent,
    SharedComponent,
    CartComponent,
    AboutComponent,
    ListingComponent,
    ListingDetailComponent,
    ImageComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    JsonpModule,
    MaterialModule,
    AppRoutingModule,
    BrowserAnimationsModule
  ],
  providers: [
    MdIconRegistry, 
    ShoppingService,
    CategoryService,
    ListingService,
    MessageService,
    ImageService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
