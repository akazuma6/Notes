import React from 'react'
import Link from 'next/link'
export default function Header() {
  return (
    <header className='flex items-center justify-between bg-amber-100 font-bold'>
      <Link href="/" className='bg-amber-400 rounded-lg'>ノート一覧</Link>
      <Link href="/post" className='bg-amber-400 rounded-lg'>メモを作成</Link>
    </header>
  )
}
