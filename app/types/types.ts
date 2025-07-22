export type BackendResponse<T> = {
  error: boolean
  message: string
  data?: T
}

export type BackendErrorResponse = {
  error: boolean
  message: string
}

export type User = {
  email: string;
  first_name: string;
  last_name: string;
};

export type Book = {
  id: number;
	title: string;
	author_id: number;
	publication_year: number;
	slug: string;
	author: Author;
	description: string;
	genres: Genre[];
	created_at: Date;
	updated_at: Date;
}

export type Author = {
	id: number;
	author_name: string;
	created_at: Date;
	updated_at: Date;
}

export type Genre = {
	id: number;
	genre_name: string;
	created_at: Date;
	updated_at: Date;
}
