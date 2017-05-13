import { Component, OnInit, Input } from '@angular/core';

import { Image } from '../shared/models/image';
import { ListingService } from '../shared/services/listing.service';

@Component({
  selector: 'app-image',
  templateUrl: './image.component.html',
  styleUrls: ['./image.component.css']
})
export class ImageComponent implements OnInit {
  @Input()
  id: string;

  image: Image[];
  constructor(private listingService: ListingService) { }

  ngOnInit() {
    this.listingService.getImages(this.id)
    .subscribe(data => {
      this.image = data;
    })
  }
}
