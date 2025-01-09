import inspect
import pathlib

class PathFinder:
    @staticmethod
    def get_base_path(the_class):
        return pathlib.Path(inspect.getfile(the_class)).resolve()
    
BASE_PATH = PathFinder.get_base_path(PathFinder).parent
DATABASE_PATH = BASE_PATH / "data/database.db"
