import { Component, OnInit } from '@angular/core';

import { Listing } from '../shared/models/listing';
import { ListingService } from '../shared/services/listing.service';
import { ShoppingService } from '../shared/services/shopping.service';
import { MessageService } from '../shared/services/message.service';

@Component({
  selector: 'app-listing',
  templateUrl: './listing.component.html',
  styleUrls: ['./listing.component.css']
})
export class ListingComponent implements OnInit {
  listings: Listing[];
  syncStatus: String;
  constructor(
    private listingService: ListingService,
    private shoppingService: ShoppingService,
    private msg: MessageService) { }

  ngOnInit() {
    this.listingService.getListings()
    .subscribe(
      data => this.listings = data,
      error => this.msg.send(error));
  }

  syncListing(id: string) {
    this.listingService.syncListing(id)
    .subscribe(data => this.syncStatus = data);
  }

  addToCart(l: Listing) {
    this.shoppingService.add(l);
  }

  listingByCart(id: number) {
    console.log(id);
  }
}
