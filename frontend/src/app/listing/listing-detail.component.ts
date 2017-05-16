import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { Listing } from '../shared/models/listing';
import { ListingService } from '../shared/services/listing.service';

@Component({
  moduleId: module.id,
  selector: 'listing-detail',
  templateUrl: 'listing-detail.component.html'
})

export class ListingDetailComponent implements OnInit {
  listing: Listing;
  constructor(private route: ActivatedRoute, private ls: ListingService) { }


  ngOnInit() { 
    this.route.params.switchMap(params => this.ls.getListing(params['id']))
    .subscribe(data => this.listing = data)
  }
}