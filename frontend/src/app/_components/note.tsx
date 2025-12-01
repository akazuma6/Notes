
import { Note } from "@/types/note";
export default function Notecomponent(props: Note) {
  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleString('ja-JP', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit'
    })
  }
  return (
    <div className="bg-white rounded-lg border border-gray-200 p-6 duration-200 mb-8 pb-4 border-b">
      {/* めもヘッダー部分 */}
      <div>
        <span>
          {formatDate(props.created_at)}
        </span>
      </div>
      {/* コンテンツ */}
      <div>
        {props.content}
      </div>
    </div>
  )
}
