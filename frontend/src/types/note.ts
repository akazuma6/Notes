export interface User {
  id: number;
  username: string;
}

export interface Note {
  id: number;
  content: string;
  created_at: string;
  user: User;
}
