import { Component, OnInit, Input } from '@angular/core';

import { Image } from '../shared/models/image';
import { ListingService } from '../shared/services/listing.service';
import { ImageService } from '../shared/services/image.service';

@Component({
  selector: 'app-image',
  templateUrl: './image.component.html',
  styleUrls: ['./image.component.css']
})
export class ImageComponent implements OnInit {
  @Input()
  id: string;

  images: Image[];
  image: Image;
  constructor(
    private listingService: ListingService,
    private imageService: ImageService) { }

  ngOnInit() {
    this.imageService.get(this.id)
    .subscribe(data => {
      this.images = data;
      this.image = data[0];
    })
  }
}
