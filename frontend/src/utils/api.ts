import { Note } from "@/types/note";
import { API_URL } from "./config";

export async function getNotes(): Promise<Note[]> {
  const response = await fetch(`${API_URL}/notes`);
  if (!response.ok) {
    throw new Error('noteの取得に失敗')
  }
  return response.json();

}


export async function createNote(content: string): Promise<Note> {
  const response = await fetch(`${API_URL}/notes`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ content }),
  }
  );
  if (!response.ok) {
    throw new Error('noteの投稿に失敗しました')
  }
  return response.json();
}
