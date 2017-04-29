export class Product {
  id: number;
  name: string;
  img: string;
  category: string;
  description: string;
  price: string;
}

export class Listing {
  listing_id: number;
  state: string;
  user_id: number;
  category_id: number;
  title: string;
  description: string;
  price: string;
  views: number;
}