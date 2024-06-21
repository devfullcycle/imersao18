export type ErrorMessageProps = {
    error: string;
  };
  
  export function ErrorMessage(props: ErrorMessageProps) {
    return (
      <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative">
        {props.error}
      </div>
    );
  }
  