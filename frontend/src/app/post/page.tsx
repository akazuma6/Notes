'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { createNote } from '@/utils/api';

export default function PostPage() {
  const [content, setContent] = useState('');
  const [loading, setLoading] = useState(false);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      setLoading(true);
      await createNote(content);
      router.push('/');
    } catch (err) {
      alert('投稿に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-8 max-w-2xl">
      {/* ヘッダ */}
      <header className="flex items-center justify-between mb-8 pb-4 border-b border-gray-200">
        <h1 className="text-3xl font-bold text-white">ノートを投稿</h1>
      </header>

      {/* フォーム: カードデザイン */}
      <form onSubmit={handleSubmit} className="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        {/* ラベル */}
        <label htmlFor="content" className="block mb-2 font-semibold text-gray-700">
          内容
        </label>

        {/* テキストエリア */}
        <textarea
          id="content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className="w-full px-4 py-3 border border-gray rounded-lg focus:ring focus:ring-blue-500 resize-none"
          rows={10}
          placeholder="ノートの内容を入力してください..."
        />

        {/* 文字数カウント */}
        <div className="text-right mt-2 text-sm text-gray-500">
          {content.length} 文字
        </div>

        {/* ボタンエリア */}
        <div className="mt-6 flex justify-end gap-3">
          {/* キャンセルボタン */}
          <Link
            href="/"
            className="px-6 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 font-semibold rounded-lg transition-colors duration-200"
          >
            キャンセル
          </Link>

          {/* 投稿ボタン */}
          <button
            type="submit"
            disabled={loading || !content.trim()}
            className="px-6 py-2 bg-blue-500 hover:bg-blue-600 disabled:bg-gray-400 disabled:cursor-not-allowed text-white font-semibold rounded-lg transition-colors duration-200"
          >
            投稿
          </button>
        </div>
      </form>
    </div>
  );
}
