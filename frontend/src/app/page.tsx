'use client'

import { useState, useEffect } from 'react';
import Link from 'next/link';
import { getNotes } from '@/utils/api';
import { Note } from '@/types/note';
import Header from './_components/header';
import Notecomponent from './_components/note';

export default function Home() {
  const [notes, setNotes] = useState<Note[]>([]);

  const fetchNotes = async () => {
    try {
      const data = await getNotes();
      setNotes(data);
    } catch (err) {
      console.log('GET NOTES失敗')
    }
  }
  // 初回レンダリング時にフェッチ
  useEffect(() => {
    fetchNotes();
  }, []);

  return (
    <>

      <div>
        {notes.length == 0 ? (
          <div>
            <p>メモがありません</p>
          </div>
        ) : (
          notes.map((note) => {
            return <Notecomponent {...note} key={note.id} />;
          })
        )}
      </div>
    </>
  )
}
