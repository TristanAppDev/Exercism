unit uLeap;

interface

type
  TYear = class
    public
      class function IsLeap(year: integer): Boolean; static;
  end;

implementation

class function TYear.IsLeap(year: integer): Boolean;
begin
  result := (year mod 4 = 0) and ((year mod 100 <> 0) or (year mod 400 = 0));
end;

end.