import React from 'react'

interface GroupHeaderProps {
    title: string,
    description: string,
    creator_id: string
}

const GroupHeader: React.FC<GroupHeaderProps> = ({
    title,
    description,
    creator_id
}) => {
  return (
    <>
        <div>
            <div>
                {title}
            </div>
            <div>
                {description}
            </div>
            <div>
                Creator Name
            </div>
        </div>
    </>
  )
}

export default GroupHeader