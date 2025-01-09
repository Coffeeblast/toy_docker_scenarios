import sqlite3
import constants
import data_model as dm

class Database():
    def __init__(self):
        db_path = constants.DATABASE_PATH
        path_exists = db_path.exists()
        if not path_exists:
            parent_folder =db_path.parent
            if not parent_folder.exists():
                parent_folder.mkdir(parents=True)
        self.__conn = sqlite3.connect(constants.DATABASE_PATH)
        if not path_exists:
            self.__bootstrap_database()

    def __bootstrap_database(self) -> None:
        sql = """
            CREATE TABLE stories (
                id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL
                , name STRING
                , author STRING
            )
        """
        curr = self.__conn.cursor()
        curr.execute(sql)
        self.__conn.commit()
        curr.close()

    def get_stories(self) -> "list[dm.Story]":
        columns = ["id", "name", "author"]
        sql = f"""
            select
                {','.join(columns)}
            from stories
        """
        curr = self.__conn.cursor()
        rows = curr.execute(sql).fetchall()
        curr.close()
        return [dm.Story(**{key: val for key, val in zip(columns, row)}) for row in rows]
    
    def add_story(self, story: dm.Story) -> None:
        sql = """
            INSERT INTO stories (name, author) VALUES (?, ?)
        """
        curr = self.__conn.cursor()
        curr.execute(sql, (story.name, story.author)).fetchall()
        self.__conn.commit()
        curr.close()

def main():
    db = Database()
    stories = [
        dm.Story(name="pernikova chalupka", author="some opilec", id=0),
        dm.Story(name="baba yaga", author="some uchylak", id=0)
    ]
    for story in stories:
        db.add_story(story)
    print(db.get_stories())



if __name__ == "__main__":
    main()